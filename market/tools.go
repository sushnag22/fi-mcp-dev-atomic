package market

type LiveDataToolInfo struct {
	Name        string
	Description string
	Endpoint    string
}

var LiveDataTools = []LiveDataToolInfo{
	{
		Name:        "Get Company Data by Name",
		Description: "Retrieve detailed financial data for a specific company using its name. The API supports full names, shortened names and common name of the stock.",
		Endpoint:    "/stock",
	},
	{
		Name:        "Industry Search",
		Description: "Search for companies within a specific industry.",
		Endpoint:    "/industry_search",
	},
	{
		Name:        "Mutual Fund Search",
		Description: "Search for mutual funds by query term.",
		Endpoint:    "/mutual_fund_search",
	},
	{
		Name:        "Trending Stocks",
		Description: "Get a snapshot of the top gaining and losing stocks at the current moment, providing valuable insights into market trends and potential trading opportunities.",
		Endpoint:    "/trending",
	},
	{
		Name:        "52 Week High Low Data",
		Description: "Retrieve data for stocks with the highest and lowest prices in the last 52 weeks from both the BSE and NSE.",
		Endpoint:    "/fetch_52_week_high_low_data",
	},
	{
		Name:        "NSE Most Active",
		Description: "Get the latest most active stocks in the National Stock Exchange (NSE) based on trading volume.",
		Endpoint:    "/NSE_most_active",
	},
	{
		Name:        "BSE Most Active",
		Description: "Get the latest most active stocks in the Bombay Stock Exchange (BSE) based on trading volume.",
		Endpoint:    "/BSE_most_active",
	},
	{
		Name:        "Mutual Funds",
		Description: "Retrieve the latest data for mutual funds, including net asset value (NAV), returns, and other details.",
		Endpoint:    "/mutual_funds",
	},
	{
		Name:        "Price Shockers",
		Description: "Get data for stocks that have experienced significant price changes in a short period of time.",
		Endpoint:    "/price_shockers",
	},
	{
		Name:        "Commodity Futures Data",
		Description: "Provides access to real-time and historical data for commodity futures contracts traded on an exchange.",
		Endpoint:    "/commodities",
	},
	{
		Name:        "Analyst Recommendations",
		Description: "Get target price information, including historical snapshots, and analyst recommendation details for a specified stock.",
		Endpoint:    "/stock_target_price",
	},
	{
		Name:        "Stock Forecasts",
		Description: "Retrieve detailed forecast information about a specified stock including EPS, Cash Flow, Capital Expenditure, and other financial metrics.",
		Endpoint:    "/stock_forecasts",
	},
	{
		Name:        "Historical Data",
		Description: "Fetch historical data for a specific stock with various time periods and filters including price, PE, market cap, etc.",
		Endpoint:    "/historical_data",
	},
	{
		Name:        "Historical Stats",
		Description: "Retrieve historical statistics for a specific stock including quarter results, balance sheet, cash flow, ratios, and shareholding patterns.",
		Endpoint:    "/historical_stats",
	},
}
