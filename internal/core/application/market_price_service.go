package application

import (
	"context"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"strconv"
	"tdex-analytics/internal/core/domain"
	"tdex-analytics/internal/core/port"
	"time"
)

const (
	fetchPriceCronExpression = ""
)

type MarketPriceService interface {
	// InsertPrice inserts market price in current moment
	InsertPrice(
		ctx context.Context,
		marketPrice MarketPrice,
	) error
	// GetPrices returns all markets prices from time in past equal to passed arg fromTime
	//if marketID is passed method will return data for all market's, otherwise only for provided one
	GetPrices(
		ctx context.Context,
		timeRange TimeRange,
		marketIDs ...string,
	) (*MarketsPrices, error)
	// StartFetchingPrices starts cron job that will periodically fetch and store prices for all markets
	StartFetchingPrices() error
}

type marketPriceService struct {
	marketPriceRepository domain.MarketPriceRepository
	marketRepository      domain.MarketRepository
	externalFetcher       port.MarketDataFetcher
	cronSvc               *cron.Cron
}

func NewMarketPriceService(
	marketPriceRepository domain.MarketPriceRepository,
	marketRepository domain.MarketRepository,
	externalFetcher port.MarketDataFetcher,
) MarketPriceService {
	return &marketPriceService{
		marketPriceRepository: marketPriceRepository,
		cronSvc:               cron.New(),
		marketRepository:      marketRepository,
		externalFetcher:       externalFetcher,
	}
}

func (m *marketPriceService) InsertPrice(
	ctx context.Context,
	marketPrice MarketPrice,
) error {
	if err := marketPrice.validate(); err != nil {
		return err
	}

	mbDomain, err := marketPrice.toDomain()
	if err != nil {
		return err
	}

	return m.marketPriceRepository.InsertPrice(ctx, *mbDomain)
}

func (m *marketPriceService) GetPrices(
	ctx context.Context,
	timeRange TimeRange,
	marketIDs ...string,
) (*MarketsPrices, error) {
	result := make(map[string][]Price)

	startTime, endTime, err := timeRange.getStartAndEndTime(time.Now())
	if err != nil {
		return nil, err
	}

	marketsPrices, err := m.marketPriceRepository.GetPricesForMarkets(
		ctx,
		startTime,
		endTime,
		marketIDs...,
	)
	if err != nil {
		return nil, err
	}

	for k, v := range marketsPrices {
		prices := make([]Price, 0)
		for _, v1 := range v {
			prices = append(prices, Price{
				BasePrice:  v1.BasePrice,
				BaseAsset:  v1.BaseAsset,
				QuotePrice: v1.QuotePrice,
				QuoteAsset: v1.QuoteAsset,
				Time:       v1.Time,
			})
		}

		result[k] = prices
	}

	return &MarketsPrices{
		MarketsPrices: result,
	}, nil
}

func (m *marketPriceService) StartFetchingPrices() error {
	if _, err := m.cronSvc.AddJob(
		fetchPriceCronExpression,
		cron.FuncJob(m.FetchPricesForAllMarkets),
	); err != nil {
		return err
	}

	return nil
}

func (m *marketPriceService) FetchPricesForAllMarkets() {
	ctx := context.Background()

	markets, err := m.marketRepository.GetAllMarkets(ctx)
	if err != nil {
		log.Errorf("FetchPricesForAllMarkets -> GetAllMarkets: %v", err)
		return
	}

	for _, v := range markets {
		go func(market domain.Market) {
			m.FetchAndInsertPrice(ctx, market)
		}(v)
	}
}

func (m *marketPriceService) FetchAndInsertPrice(
	ctx context.Context,
	market domain.Market,
) {
	price, err := m.externalFetcher.FetchPrice(ctx, market)
	if err != nil {
		log.Errorf("FetchAndInsertPrice -> FetchPrice: %v", err)
		return
	}

	if err := m.InsertPrice(ctx, MarketPrice{
		MarketID:   strconv.Itoa(market.ID),
		BasePrice:  price.BasePrice,
		BaseAsset:  market.BaseAsset,
		QuotePrice: price.QuotePrice,
		QuoteAsset: market.QuoteAsset,
		Time:       time.Now(),
	}); err != nil {
		log.Errorf("FetchAndInsertPrice -> InsertPrice: %v", err)
		return
	}
}
