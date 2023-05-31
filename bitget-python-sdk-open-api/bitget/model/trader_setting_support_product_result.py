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


class TraderSettingSupportProductResult(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        
        class properties:
            openCopyTrace = schemas.BoolSchema
            productCode = schemas.StrSchema
            productName = schemas.StrSchema
            __annotations__ = {
                "openCopyTrace": openCopyTrace,
                "productCode": productCode,
                "productName": productName,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["openCopyTrace"]) -> MetaOapg.properties.openCopyTrace: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["productCode"]) -> MetaOapg.properties.productCode: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["productName"]) -> MetaOapg.properties.productName: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["openCopyTrace", "productCode", "productName", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["openCopyTrace"]) -> typing.Union[MetaOapg.properties.openCopyTrace, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["productCode"]) -> typing.Union[MetaOapg.properties.productCode, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["productName"]) -> typing.Union[MetaOapg.properties.productName, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["openCopyTrace", "productCode", "productName", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        openCopyTrace: typing.Union[MetaOapg.properties.openCopyTrace, bool, schemas.Unset] = schemas.unset,
        productCode: typing.Union[MetaOapg.properties.productCode, str, schemas.Unset] = schemas.unset,
        productName: typing.Union[MetaOapg.properties.productName, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'TraderSettingSupportProductResult':
        return super().__new__(
            cls,
            *args,
            openCopyTrace=openCopyTrace,
            productCode=productCode,
            productName=productName,
            _configuration=_configuration,
            **kwargs,
        )
