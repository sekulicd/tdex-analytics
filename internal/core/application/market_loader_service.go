package application

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tdex-network/tdex-analytics/internal/core/domain"
	tdexmarketloader "github.com/tdex-network/tdex-analytics/pkg/tdex-market-loader"
	"time"
)

const (
	//cron expression for fetching markets, every hour
	fetchMarketsCronExpression = "0 * * * *"
)

type MarketsLoaderService interface {
	StartFetchingMarketsJob() error
}

type marketsLoaderService struct {
	marketRepository    domain.MarketRepository
	tdexMarketLoaderSvc tdexmarketloader.Service
	cronSvc             *cron.Cron
}

func NewMarketsLoaderService(
	marketRepository domain.MarketRepository,
	tdexMarketLoaderSvc tdexmarketloader.Service,
) MarketsLoaderService {
	return &marketsLoaderService{
		marketRepository:    marketRepository,
		tdexMarketLoaderSvc: tdexMarketLoaderSvc,
		cronSvc:             cron.New(),
	}
}

func (m *marketsLoaderService) StartFetchingMarketsJob() error {
	// run initially
	go m.FetchMarkets()

	if _, err := m.cronSvc.AddJob(
		fetchMarketsCronExpression,
		cron.FuncJob(m.FetchMarkets),
	); err != nil {
		return err
	}

	m.cronSvc.Start()

	return nil
}

func (m *marketsLoaderService) FetchMarkets() {
	log.Infof("job FetchMarkets at: %v", time.Now())
	//TODO add context with timeout
	liquidityProviders, err := m.tdexMarketLoaderSvc.FetchProvidersMarkets(
		context.Background(),
	)
	if err != nil {
		log.Errorf("FetchMarkets -> FetchProvidersMarkets: %v", err)
		return
	}

	//markets already stored in db
	existingMarkets, err := m.marketRepository.GetAllMarkets(context.Background())
	if err != nil {
		log.Errorf("FetchMarkets -> GetAllMarkets: %v", err)
		return
	}

	//assume that external service, from which we are fetching markets, should return active markets
	activeMarkets := make([]domain.Market, 0)
	for _, v := range liquidityProviders {
		for _, market := range v.Markets {
			activeMarkets = append(activeMarkets, domain.Market{
				ProviderName: v.Name,
				Url:          v.Endpoint,
				BaseAsset:    market.BaseAsset,
				QuoteAsset:   market.QuoteAsset,
				Active:       true,
			})
		}
	}

	if err := m.updateMarketActiveStatusAndInsertNew(activeMarkets, existingMarkets); err != nil {
		log.Errorf("FetchMarkets -> updateMarketActiveStatusAndInsertNew: %v", err)
	}
}

func (m *marketsLoaderService) updateMarketActiveStatusAndInsertNew(
	existingMarkets []domain.Market,
	activeMarkets []domain.Market,
) error {
	//since markets can go on and off, we need to update markets accordingly
	//3 cases here:
	//1. existing market is not in active markets -> set market as inactive
	//2. existing market is in active markets -> set market as active
	//3. active market is not in existing markets -> create new market
	marketsNotInActiveList := make([]domain.Market, 0) //to be inactivated
	marketsInActiveList := make([]domain.Market, 0)    //to be activated
	for _, v1 := range existingMarkets {
		found := false
		for _, v2 := range activeMarkets {
			if isSameMarket(v1, v2) {
				found = true
				marketsInActiveList = append(marketsInActiveList, v1)
			}
		}

		if !found {
			marketsNotInActiveList = append(marketsNotInActiveList, v1)
		}
	}

	if len(marketsNotInActiveList) > 0 {
		for _, v := range marketsNotInActiveList {
			if err := m.marketRepository.InactivateMarket(
				context.Background(),
				v.ID, //TODO does id exist
			); err != nil {
				return fmt.Errorf("updateMarketActiveStatusAndInsertNew -> InactivateMarket: %v", err)
			}
		}
	}

	if len(marketsInActiveList) > 0 {
		for _, v := range marketsInActiveList {
			if err := m.marketRepository.ActivateMarket(
				context.Background(),
				v.ID, //TODO does id exist
			); err != nil {
				return fmt.Errorf("updateMarketActiveStatusAndInsertNew -> ActivateMarket: %v", err)
			}
		}
	}

	//now we need to add new markets to db, if market already exist, insert wont do anything
	for _, v := range activeMarkets {
		if err := m.marketRepository.InsertMarket(context.Background(), domain.Market{
			ProviderName: v.ProviderName,
			Url:          v.Url,
			BaseAsset:    v.BaseAsset,
			QuoteAsset:   v.QuoteAsset,
			Active:       true,
		}); err != nil {
			return fmt.Errorf("updateMarketActiveStatusAndInsertNew -> InsertMarket: %v", err)
		}
	}

	return nil
}

func isSameMarket(m1, m2 domain.Market) bool {
	return m1.Url == m2.Url &&
		m1.BaseAsset == m2.BaseAsset &&
		m1.QuoteAsset == m2.QuoteAsset
}
