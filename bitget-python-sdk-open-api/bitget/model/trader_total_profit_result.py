# coding: utf-8

"""
    Bitget Open API

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 2.0.0
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from bitget import schemas  # noqa: F401


class TraderTotalProfitResult(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        
        class properties:
            sumProfit = schemas.StrSchema
            waitProfit = schemas.StrSchema
            yesterdaySplitProfit = schemas.StrSchema
            yesterdayTimeStamp = schemas.StrSchema
            __annotations__ = {
                "sumProfit": sumProfit,
                "waitProfit": waitProfit,
                "yesterdaySplitProfit": yesterdaySplitProfit,
                "yesterdayTimeStamp": yesterdayTimeStamp,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["sumProfit"]) -> MetaOapg.properties.sumProfit: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["waitProfit"]) -> MetaOapg.properties.waitProfit: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["yesterdaySplitProfit"]) -> MetaOapg.properties.yesterdaySplitProfit: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["yesterdayTimeStamp"]) -> MetaOapg.properties.yesterdayTimeStamp: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["sumProfit", "waitProfit", "yesterdaySplitProfit", "yesterdayTimeStamp", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["sumProfit"]) -> typing.Union[MetaOapg.properties.sumProfit, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["waitProfit"]) -> typing.Union[MetaOapg.properties.waitProfit, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["yesterdaySplitProfit"]) -> typing.Union[MetaOapg.properties.yesterdaySplitProfit, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["yesterdayTimeStamp"]) -> typing.Union[MetaOapg.properties.yesterdayTimeStamp, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["sumProfit", "waitProfit", "yesterdaySplitProfit", "yesterdayTimeStamp", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        sumProfit: typing.Union[MetaOapg.properties.sumProfit, str, schemas.Unset] = schemas.unset,
        waitProfit: typing.Union[MetaOapg.properties.waitProfit, str, schemas.Unset] = schemas.unset,
        yesterdaySplitProfit: typing.Union[MetaOapg.properties.yesterdaySplitProfit, str, schemas.Unset] = schemas.unset,
        yesterdayTimeStamp: typing.Union[MetaOapg.properties.yesterdayTimeStamp, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'TraderTotalProfitResult':
        return super().__new__(
            cls,
            *args,
            sumProfit=sumProfit,
            waitProfit=waitProfit,
            yesterdaySplitProfit=yesterdaySplitProfit,
            yesterdayTimeStamp=yesterdayTimeStamp,
            _configuration=_configuration,
            **kwargs,
        )
