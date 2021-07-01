import bitget.spot.public_api as public
import bitget.spot.market_api as market
import bitget.spot.account_api as account
import bitget.spot.order_api as order

if __name__ == '__main__':
    api_key = ""
    secret_key = ""
    passphrase = ""  # 口令

    symbol = 'bftusdt_spbl'

    # spot 获取币种信息
    publicApi = public.PublicApi(api_key, secret_key, passphrase, use_server_time=True, first=False);
    # result = publicApi.times()
    # print(result)

    # spot 获取交易对信息
    # result = publicApi.products()
    # print(result)

    # spot 获取单个交易对信息
    result = publicApi.product('bftusdt_spbl')
    print(result)



    # marketApi = market.MarketApi(api_key, secret_key, passphrase, use_server_time=False, first=False);
    # result = marketApi.fills(symbol, limit=50)
    # print(result)

    # result = marketApi.depth(symbol, limit=50, type='step0')
    # print(result)

    # result = marketApi.ticker(symbol)
    # print(result)

    # result = marketApi.tickers()
    # print(result)

    # result = marketApi.candles(symbol, period='1min', after='1624352586', before='1624356186', limit=100)
    # print(result)

    accountApi = account.AccountApi(api_key, secret_key, passphrase, use_server_time=False, first=False)

    # result = accountApi.assets()
    # print(result)

    # result = accountApi.bills(coinId='2', groupType='transaction', bizType='buy', limit=10)
    # print(result)

    # orderApi = order.OrderApi(api_key, secret_key, passphrase, use_server_time=False, first=False)
    # result = orderApi.orders(symbol='bftusdt_spbl', quantity='20', side='buy', orderType='market', force='normal', clientOrderId='spot#2392818kk')
    # print(result)

    # order_data=[{"price":"2.30222","quantity":"1","side":"buy","orderType":"limit","force":"normal","client_oid":"spot#jidhuu19399"}, {"price":"2.30111","quantity":"1","side":"buy","orderType":"limit","force":"normal","client_oid":"spot#akncnai8821"}]
    # result = orderApi.batch_orders(symbol='bftusdt_spbl', order_data=order_data)
    # print(result)

    # result = orderApi.order_info(symbol='bftusdt_spbl')
    # print(result)

    # result = orderApi.cancel_orders(symbol='bftusdt_spbl', orderId='791171749756964864')
    # print(result)

    # result = orderApi.cancel_batch_orders(symbol='bftusdt_spbl', orderIds=['793726305732825088', '793726305804128256'])
    # print(result)

    # result = orderApi.open_order(symbol='bftusdt_spbl')
    # print(result)

    # result = orderApi.history(symbol='bftusdt_spbl')
    # print(result)

    # result = orderApi.fills(symbol='bftusdt_spbl', orderId='791113184589549568')
    # print(result)