# Table of Contents

- [v0.10.0](#v0100---20191027)
- [v0.9.0](#v090---20190824)
- [v0.8.0](#v080---20190821)
- [v0.7.0](#v070---20190813)
- [v0.6.2](#v062---20190809)
- [v0.6.1](#v061---20190809)
- [v0.6.0](#v060---20190809)
- [v0.5.1](#v051---20190805)
- [v0.5.0](#v050---20190607)
- [v0.4.1](#v041---20190411)
- [v0.4.0](#v040---20190406)
- [0.3.0](#030---20181219)
- [0.2.0](#020---20181219)
- [0.1.0](#010---20181201)

## [v0.10.0] - 2019/10/27

### Summary

- This release adds support for Kong 1.4.

### Added

- `HostHeader` field has been added to Upstream struct.
- `Tags` field has been added to the following types:
  - `KeyAuth`
  - `Basicauth`
  - `HMACAuth`
  - `Oauth2Credential`
  - `ACLGroup`
  - `JWTAuth`

## [v0.9.0] - 2019/08/24

### Breaking changes

- `client.Do()` returns a response object even on errors so that
  clients can inspect the response directly when needed.
  The error condition has changed to include HTTP status codes 300 to 399
  as success and not failure.
  [b1f9eda31](https://github.com/hbagdi/go-kong/commit/b1f9eda311e1d9c9d6b0f5a5e33a3d399d85faf6)
- `String()` method has been dropped from all types defined in this package.

### Added

- `NewRequest()` method helping with creating HTTP requests is now exported
- URLs parsed inside the package are more robust.
- New method `GetByCustomID` has been introduced to fetch Consumers by
  `custom_id`.

## [v0.8.0] - 2019/08/21

### Added

- Oauth2Credential type and service has been added
  which can be used to create Oauth2 credentials in Kong for some
  Oauth2 flows.

## [v0.7.0] - 2019/08/13

### Summary

This release brings support for CRUD methods for
authentication credentials in Kong.

### Added

- The following credentials and corresponding services have been added:
  - `key-auth`
  - `basic-auth`
  - `hmac-auth`
  - `jwt`
  - `acl`

## [v0.6.2] - 2019/08/09

### Fix

- Add missing omitempty tag to ClientCertificate field

## [v0.6.1] - 2019/08/09

### Fix

- Fix a typo in Service struct definition for YAML tag

## [v0.6.0] - 2019/08/09

### Summary

- This release adds support for Kong 1.3.

### Breaking change

- `Validator` Interface has been dropped and Valid() method from
  all entities is dropped.
  [#8](https://github.com/hbagdi/go-kong/pull/8)

### Added

- `Headers` field has been added to Route struct.
  [#5](https://github.com/hbagdi/go-kong/pull/5)
- `ClientCertificate` field has been added to Service struct.
  [#5](https://github.com/hbagdi/go-kong/pull/5)
- `CACertificate` is a new core entity in Kong. A struct to represent
  it and a corresponding new service is added.
  [#5](https://github.com/hbagdi/go-kong/pull/5)
- `Algorithm` field has been added to Upstream struct.
  [#9](https://github.com/hbagdi/go-kong/pull/9)

## [v0.5.1] - 2019/08/05

### Fix

- Add missing healthchecks.active.unhealthy.interval field to Upstream
  [#6](https://github.com/hbagdi/go-kong/issues/6)

## [v0.5.0] - 2019/06/07

### Summary

- This release adds support for Kong 1.2.

### Added

- Added HTTPSRedirectStatusCode property to Route struct.
  [#3](https://github.com/hbagdi/go-kong/pull/3)

### Breaking change

- `Create()` for Custom Entities now supports HTTP PUT method.
  If `id` is specified in the object, it will be used to PUT the entity.
  This was always POST previously.
  [#3](https://github.com/hbagdi/go-kong/pull/3)

## [v0.4.1] - 2019/04/11

### Fix

- Add `omitempty` property to Upstream fields for Kong 1.0 compatibility

## [v0.4.0] - 2019/04/06

### Summary

- This release adds support for features released in Kong 1.1.
  This version is compatible with Kong 1.0 and Kong 1.1.

### Breaking Change

- Please note that the version naming scheme for this library has changed from
  `x.y.z` to `vX.Y.Z`. This is to ensure compatibility with Go modules.

### Added

- `Tags` field has been added to all Kong Core entity structs.
- List methods now support tag based filtering introduced in Kong 1.1.
  Tags can be ANDed or ORed together. `ListOpt` struct can be used to
  specify the tags for filtering.
- `Protocols` field has been added to Plugin struct.
- New fields `Type`, `HTTPSSni` and `HTTPSVerifyCertificate` have been
  introduced for Active HTTPS healthchecks.
- `TargetService` has two new methods `MarkHealthy()` and `MarkUnhealthy()`
  to change the health of a target.

## [0.3.0] - 2018/12/19

### Summary

- This release adds support for Kong 1.0.
  It is not compatible with 0.x.y  versions of Kong due to breaking
  Admin API changes as the deprecated API entity is dropped.
- The code and API for the library is same as 0.2.0, with the exception
  that struct defs and services related to `API` is dropped.

### Breaking changes

- `API` struct definition is no longer available.
- `APIService` is no longer available. Please ensure your code doesn't rely
  on these before upgrading.
- `Plugin` struct has dropped the `API` field.

## [0.2.0] - 2018/12/19

### Summary

- This release adds support for Kong 0.15.x.
  It is not compatible with any other versions of Kong due to breaking
  Admin API changes in Kong for Plugins, Upstreams and Targets entities.

### Breaking changes

- `Target` struct now has an `Upstream` member in place of `UpstreamID`.
- `Plugin` struct now has `Consumer`, `API`, `Route`, `Service` members
  instead of `ConsumerID`, `APIID`, `RouteID` and `ServiceID`.

### Added

- `RunOn` property has been added to `Plugin`.
- New properties are added to `Route` for L4 proxy support.

## [0.1.0] - 2018/12/01

### Summary

- Debut release of this library
- This release comes with support for Kong 0.14.x
- The library is not expected to work with previous or later
  releases of Kong since every release of Kong is introducing breaking changes
  to the Admin API.

[v0.10.0]: https://github.com/hbagdi/go-kong/compare/v0.9.0...v0.10.0
[v0.9.0]: https://github.com/hbagdi/go-kong/compare/v0.8.0...v0.9.0
[v0.8.0]: https://github.com/hbagdi/go-kong/compare/v0.7.0...v0.8.0
[v0.7.0]: https://github.com/hbagdi/go-kong/compare/v0.6.2...v0.7.0
[v0.6.2]: https://github.com/hbagdi/go-kong/compare/v0.6.1...v0.6.2
[v0.6.1]: https://github.com/hbagdi/go-kong/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/hbagdi/go-kong/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/hbagdi/go-kong/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/hbagdi/go-kong/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/hbagdi/go-kong/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/hbagdi/go-kong/compare/0.3.0...v0.4.0
[0.3.0]: https://github.com/hbagdi/go-kong/compare/0.2.0...0.3.0
[0.2.0]: https://github.com/hbagdi/go-kong/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/hbagdi/go-kong/compare/87666c7fe73477d1874d35d690301241cd23059f...0.1.0
