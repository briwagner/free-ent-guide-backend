# Free Entertainment Guide

This project is a 100% Go backend to support a Javascript-driven frontend web application. The backend manages the following:
* user account creation and authentication
* integration with <a href="https://github.com/shaj13/go-guardian">Go-Guardian</a> for authorization via basic auth and JWT
* proxy and caching for requests to third-party data providers

## Authentication and Authorization

The Go-Guardian library allows us to register multiple methods for authenticating requests. FEG supports basic auth, which integrates with a database. If approved, the client receives a JWT which is stored by the frontend application and used for subsequent requests.

## Proxy and Caching

The Go server is an efficient method for proxying requests to third-party data providers. Details about URL paths and authentication credentials for the third-party providers are not exposed to the client on the frontend, a common problem for Javascript-only applications.

By integrating with database storage, we can also cache responses for reuse instead of forwarding all requests to the third-party provider.

## CLI

The cli component helps to manage admin tasks against the database. This process uses Docker to run a Python image that pulls data from external feeds.

## Sqlc

cd models
sqlc generate

Is there a better way to run this?