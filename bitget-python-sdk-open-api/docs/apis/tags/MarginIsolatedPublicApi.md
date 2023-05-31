<a name="__pageTop"></a>
# bitget.apis.tags.margin_isolated_public_api.MarginIsolatedPublicApi

All URIs are relative to *https://api.bitget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**margin_isolated_public_interest_rate_and_limit**](#margin_isolated_public_interest_rate_and_limit) | **get** /api/margin/v1/isolated/public/interestRateAndLimit | interestRateAndLimit
[**margin_isolated_public_tier_data**](#margin_isolated_public_tier_data) | **get** /api/margin/v1/isolated/public/tierData | tierData

# **margin_isolated_public_interest_rate_and_limit**
<a name="margin_isolated_public_interest_rate_and_limit"></a>
> ApiResponseResultOfListOfMarginIsolatedRateAndLimitResult margin_isolated_public_interest_rate_and_limit(symbol)

interestRateAndLimit

Get InterestRateAndLimit

### Example

```python
import bitget
from bitget.apis.tags import margin_isolated_public_api
from bitget.model.api_response_result_of_void import ApiResponseResultOfVoid
from bitget.model.api_response_result_of_list_of_margin_isolated_rate_and_limit_result import ApiResponseResultOfListOfMarginIsolatedRateAndLimitResult
from pprint import pprint
# Defining the host is optional and defaults to https://api.bitget.com
# See configuration.py for a list of all supported configuration parameters.
configuration = bitget.Configuration(
    host = "https://api.bitget.com"
)

# Enter a context with an instance of the API client
with bitget.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = margin_isolated_public_api.MarginIsolatedPublicApi(api_client)

    # example passing only required values which don't have defaults set
    query_params = {
        'symbol': "BTCUSDT",
    }
    try:
        # interestRateAndLimit
        api_response = api_instance.margin_isolated_public_interest_rate_and_limit(
            query_params=query_params,
        )
        pprint(api_response)
    except bitget.ApiException as e:
        print("Exception when calling MarginIsolatedPublicApi->margin_isolated_public_interest_rate_and_limit: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
symbol | SymbolSchema | | 


# SymbolSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#margin_isolated_public_interest_rate_and_limit.ApiResponseFor200) | OK
400 | [ApiResponseFor400](#margin_isolated_public_interest_rate_and_limit.ApiResponseFor400) | Bad Request
429 | [ApiResponseFor429](#margin_isolated_public_interest_rate_and_limit.ApiResponseFor429) | Gateway Frequency Limit
500 | [ApiResponseFor500](#margin_isolated_public_interest_rate_and_limit.ApiResponseFor500) | Server Error

#### margin_isolated_public_interest_rate_and_limit.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfListOfMarginIsolatedRateAndLimitResult**](../../models/ApiResponseResultOfListOfMarginIsolatedRateAndLimitResult.md) |  | 


#### margin_isolated_public_interest_rate_and_limit.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor400ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor400ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfVoid**](../../models/ApiResponseResultOfVoid.md) |  | 


#### margin_isolated_public_interest_rate_and_limit.ApiResponseFor429
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor429ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor429ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfVoid**](../../models/ApiResponseResultOfVoid.md) |  | 


#### margin_isolated_public_interest_rate_and_limit.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor500ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor500ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfVoid**](../../models/ApiResponseResultOfVoid.md) |  | 


### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **margin_isolated_public_tier_data**
<a name="margin_isolated_public_tier_data"></a>
> ApiResponseResultOfListOfMarginIsolatedLevelResult margin_isolated_public_tier_data(symbol)

tierData

Get TierData

### Example

```python
import bitget
from bitget.apis.tags import margin_isolated_public_api
from bitget.model.api_response_result_of_void import ApiResponseResultOfVoid
from bitget.model.api_response_result_of_list_of_margin_isolated_level_result import ApiResponseResultOfListOfMarginIsolatedLevelResult
from pprint import pprint
# Defining the host is optional and defaults to https://api.bitget.com
# See configuration.py for a list of all supported configuration parameters.
configuration = bitget.Configuration(
    host = "https://api.bitget.com"
)

# Enter a context with an instance of the API client
with bitget.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = margin_isolated_public_api.MarginIsolatedPublicApi(api_client)

    # example passing only required values which don't have defaults set
    query_params = {
        'symbol': "BTCUSDT",
    }
    try:
        # tierData
        api_response = api_instance.margin_isolated_public_tier_data(
            query_params=query_params,
        )
        pprint(api_response)
    except bitget.ApiException as e:
        print("Exception when calling MarginIsolatedPublicApi->margin_isolated_public_tier_data: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
symbol | SymbolSchema | | 


# SymbolSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#margin_isolated_public_tier_data.ApiResponseFor200) | OK
400 | [ApiResponseFor400](#margin_isolated_public_tier_data.ApiResponseFor400) | Bad Request
429 | [ApiResponseFor429](#margin_isolated_public_tier_data.ApiResponseFor429) | Gateway Frequency Limit
500 | [ApiResponseFor500](#margin_isolated_public_tier_data.ApiResponseFor500) | Server Error

#### margin_isolated_public_tier_data.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfListOfMarginIsolatedLevelResult**](../../models/ApiResponseResultOfListOfMarginIsolatedLevelResult.md) |  | 


#### margin_isolated_public_tier_data.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor400ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor400ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfVoid**](../../models/ApiResponseResultOfVoid.md) |  | 


#### margin_isolated_public_tier_data.ApiResponseFor429
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor429ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor429ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfVoid**](../../models/ApiResponseResultOfVoid.md) |  | 


#### margin_isolated_public_tier_data.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor500ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor500ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponseResultOfVoid**](../../models/ApiResponseResultOfVoid.md) |  | 


### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)
