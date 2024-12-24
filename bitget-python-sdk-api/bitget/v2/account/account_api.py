#!/usr/bin/python
from bitget.client import Client
from bitget.consts import GET


class AccountApi(Client):
    def __init__(self, api_key, api_secret_key, passphrase, use_server_time=False):
        Client.__init__(self, api_key, api_secret_key, passphrase, use_server_time)

    def all_account_balance(self):
        return self._request_with_params(GET, "/api/v2/account/all-account-balance", {})