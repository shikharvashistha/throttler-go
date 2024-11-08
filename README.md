# throttler-go

## Description

This is a simple throttler for Go. It is based on the [throttler](https://www.django-rest-framework.org/api-guide/throttling/) from Django REST Framework.

Throttling is a way to limit the number of requests a client can make in a given period of time. This is useful to prevent abuse of an API and to protect the server from being overloaded.

## How throttling is determined

Throttling is determined by the following factors:

- The client's IP address.
- Custom

## API reference

### `AnonRateThrottler`

This rate throttler is used to throttle anonymous requests and is based on the IP address of the client. The IP address is used to generate a cache key to throttle the client.


### `CustomRateThrottler`

This rate throttler can be used to throttle requests based on a custom key. The key is generated by the `getCacheKey` method. The default implementation of this method uses the IP address of the client. `getCacheKey` can be overridden to use a different key. It has `getCacheKey func(r *http.Request) (string, error))` as a parameter. The `getCacheKey` function should return a string that will be used as the cache key. If the `getCacheKey` function returns an error, the request will be throttled.


## Setup

### Setting up the cache/customizing the cache key

If you want to use a cache other than the default one, you can modify/implement [interface.go](pkg/store/interface.go) to use your own cache.

### Setting up [GetIndent](pkg/utils/get_indent.go)

The `GetIndent` function can also be modified to use a different key. It has `GetIndent func(r *http.Request, numProxies int) string` as a parameter. The `GetIndent` function should return a string that will be used as the cache key. So if you want to use a different key, you can modify this function as per requirement.


### Setting up the AnonRateThrottler

The `AnonRateThrottler` is used to throttle anonymous requests. It is based on the IP address of the client. The IP address is used to generate a cache key to throttle the client. The `AnonRateThrottler` takes the following parameters:

- `rate`: The number of requests allowed in the given `duration`.
- `duration`: The duration of the rate limit.
- `scope`: The scope of the rate limit. This is used to generate the cache key.
- `numProxies`: The number of proxies in front of the server. This is used to get the client's IP address.
- `kvs`: The key-value store used to store the cache.

If you want to use a different cache, you can implement [interface.go](pkg/store/interface.go) to use your own cache.

### Setting up the CustomRateThrottler

The `CustomRateThrottler` can be used to throttle requests based on a custom key. The key is generated by the `getCacheKey` method. The default implementation of this method uses the IP address of the client. `getCacheKey` can be overridden to use a different key. It has `getCacheKey func(r *http.Request) (string, error))` as a parameter. The `getCacheKey` function should return a string that will be used as the cache key. If the `getCacheKey` function returns an error, the request will be throttled.

The `CustomRateThrottler` takes the following parameters:

- `reqAllowed`: The number of requests allowed in the given `duration`.
- `inDur`: The duration of the rate limit.
- `scope`: The scope of the rate limit. This is used to generate the cache key.
- `kvs`: The key-value store used to store the cache.
- `getCacheKey`: The function used to generate the cache key.


## Examples

### AnonRateThrottler

`anonymous_throttle := middleware.GetAnonymousThrottle(a, b, "c", d, e)`

- a: The number of requests allowed in the given `duration`.
- b: The duration of the rate limit.
- c: The scope of the rate limit. This is used to generate the cache key.
- d: The number of proxies in front of the server. This is used to get the client's IP address.
- e: The key-value store used to store the cache.


The following code can be used to check if the request is allowed and to get the time to wait before the next request can be made.

`c, _ := anonymous_throttle.AllowRequest(r)`
`wait, _ := anonymous_throttle.Wait()`


### CustomRateThrottler


`custom_throttle := middleware.GetCustomThrottle(a, b, "c", d, e)`

- a: The number of requests allowed in the given `duration`.
- b: The duration of the rate limit.
- c: The scope of the rate limit. This is used to generate the cache key.
- d: The key-value store used to store the cache.
- e: The function used to generate the cache key.

The following code can be used to check if the request is allowed and to get the time to wait before the next request can be made.

`c, _ := custom_throttle.AllowRequest(r)`
`wait, _ := custom_throttle.Wait()`