# Changelog

## [2.0.0](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/v1.6.0...v2.0.0) (2026-05-08)


### ⚠ BREAKING CHANGES

* GH_APP_ID is replaced by GH_APP_CLIENT_ID for GitHub App authentication.

### Features

* use GitHub App client ID for authentication ([#48](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/48)) ([351436e](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/351436e9fc21aa9d583ab762b271c95d85e13170))


### Bug Fixes

* **deps:** update module github.com/jferrl/go-githubauth to v1.6.0 ([#46](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/46)) ([b962bec](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b962bec68956edecab52a1dcff3c4e5e7a887279))
* request GitHub App tokens with empty permissions ([#44](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/44)) ([069b446](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/069b446d3f7e0eaebd198a91fdbb2a8947088563))

## [1.6.0](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/v1.5.0...v1.6.0) (2026-05-07)


### Features

* add missing rate limit metrics ([#42](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/42)) ([2da56dd](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/2da56ddfe96b489f6cc9c3ae578ab36e16af1a6a))

## [1.5.0](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/v1.4.2...v1.5.0) (2026-05-07)


### Features

* add healthcheck endpoint ([#39](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/39)) ([9c46b70](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9c46b70415009ae5e257a91c2ee2f8e858a4bc89))

## [1.4.2](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/v1.4.1...v1.4.2) (2026-05-07)


### Bug Fixes

* **deps:** update module github.com/google/go-github/v61 to v85 ([#25](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/25)) ([7ddc59c](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/7ddc59c90075a2952d8fa2d7560012d85902070f))

## [1.4.1](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/v1.4.0...v1.4.1) (2026-05-07)


### Bug Fixes

* goreleaser ([#35](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/35)) ([db399d1](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/db399d1ebd8bee47c39a62417dfc1469c6fc1736))

## [1.4.0](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/v1.3.0...v1.4.0) (2026-05-07)


### Features

* add golangci-lint. ([9a057bd](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9a057bdb4a067d310f2d792185c8434d7f9369e0))
* add lint. ([61f21e4](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/61f21e45a0af555396b1e40a3f60eafe5be7c5c4))
* devcontainer wip ([b2b22ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b2b22ffc7839e95a9a7beb737a04fda65d6a5705))
* impl release flow and e2e test ([a750c6d](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/a750c6dd30dd33afc3c82e26cb441f587712e822))
* impl release flow and e2e test ([c4b54fb](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/c4b54fb46c40d6248e2599739a98a65f8ca6b7ca))
* run lint, test, e2e-test in parallel ([9ebc7f2](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9ebc7f22d1333f217842fcd62a72a4acb7f811ba))


### Bug Fixes

* **deps:** update module github.com/rs/zerolog to v1.35.1 ([#16](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/16)) ([00c3915](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/00c391578b42d0a91229a92a31b5a9cf1d9b1162))
* **deps:** update module go.uber.org/dig to v1.19.0 ([#18](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/18)) ([53082ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/53082ff4b51d61dc15b38ab64c3959dd9bebb653))
* fix a test case name. ([b4f49c3](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b4f49c333135f60d5eda7b00006dc3e97e72c51d))
* install moq on release ([#13](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/13)) ([04da01c](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/04da01c77436b5c072f0bce964a90960e829254d))
* publish as repository owner ([#10](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/10)) ([45d71de](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/45d71de722cb11539d474a6050308638764fc204))
* release ([#29](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/29)) ([5c111e7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/5c111e714b25b9b8d08782d5313067ff6e82564a))
* release ([#34](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/34)) ([6b70350](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/6b70350a03a015763bce89b4c03e8dd317f5637a))
* release-please ([#32](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/32)) ([ced0d38](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/ced0d387c7b9a5e627fa181289eabc5c1e27c477))
* renovate vonfig ([#14](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/14)) ([e7335ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/e7335ffa1e206102c15518826980770e29ec328a))
* run ci on release-please PRs ([#7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/7)) ([1974317](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/19743170ccf5fd7985a80d0a76177886ecaa789a))
* update GHA workflow to wait for all test jobs. ([6592646](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/65926463eb49c2b1100c893a1b5974fe3f29d9a9))

## [1.3.0](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/v1.2.0...v1.3.0) (2026-05-07)


### Features

* add golangci-lint. ([9a057bd](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9a057bdb4a067d310f2d792185c8434d7f9369e0))
* add lint. ([61f21e4](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/61f21e45a0af555396b1e40a3f60eafe5be7c5c4))
* devcontainer wip ([b2b22ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b2b22ffc7839e95a9a7beb737a04fda65d6a5705))
* impl release flow and e2e test ([a750c6d](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/a750c6dd30dd33afc3c82e26cb441f587712e822))
* impl release flow and e2e test ([c4b54fb](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/c4b54fb46c40d6248e2599739a98a65f8ca6b7ca))
* run lint, test, e2e-test in parallel ([9ebc7f2](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9ebc7f22d1333f217842fcd62a72a4acb7f811ba))


### Bug Fixes

* **deps:** update module github.com/rs/zerolog to v1.35.1 ([#16](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/16)) ([00c3915](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/00c391578b42d0a91229a92a31b5a9cf1d9b1162))
* **deps:** update module go.uber.org/dig to v1.19.0 ([#18](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/18)) ([53082ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/53082ff4b51d61dc15b38ab64c3959dd9bebb653))
* fix a test case name. ([b4f49c3](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b4f49c333135f60d5eda7b00006dc3e97e72c51d))
* install moq on release ([#13](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/13)) ([04da01c](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/04da01c77436b5c072f0bce964a90960e829254d))
* publish as repository owner ([#10](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/10)) ([45d71de](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/45d71de722cb11539d474a6050308638764fc204))
* release ([#29](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/29)) ([5c111e7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/5c111e714b25b9b8d08782d5313067ff6e82564a))
* release-please ([#32](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/32)) ([ced0d38](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/ced0d387c7b9a5e627fa181289eabc5c1e27c477))
* renovate vonfig ([#14](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/14)) ([e7335ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/e7335ffa1e206102c15518826980770e29ec328a))
* run ci on release-please PRs ([#7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/7)) ([1974317](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/19743170ccf5fd7985a80d0a76177886ecaa789a))
* update GHA workflow to wait for all test jobs. ([6592646](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/65926463eb49c2b1100c893a1b5974fe3f29d9a9))

## [1.2.0](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/gh-ratelimit-metrics-exporter-v1.1.0...gh-ratelimit-metrics-exporter-v1.2.0) (2026-05-07)


### Features

* add golangci-lint. ([9a057bd](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9a057bdb4a067d310f2d792185c8434d7f9369e0))
* add lint. ([61f21e4](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/61f21e45a0af555396b1e40a3f60eafe5be7c5c4))
* devcontainer wip ([b2b22ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b2b22ffc7839e95a9a7beb737a04fda65d6a5705))
* impl release flow and e2e test ([a750c6d](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/a750c6dd30dd33afc3c82e26cb441f587712e822))
* impl release flow and e2e test ([c4b54fb](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/c4b54fb46c40d6248e2599739a98a65f8ca6b7ca))
* run lint, test, e2e-test in parallel ([9ebc7f2](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9ebc7f22d1333f217842fcd62a72a4acb7f811ba))


### Bug Fixes

* **deps:** update module github.com/rs/zerolog to v1.35.1 ([#16](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/16)) ([00c3915](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/00c391578b42d0a91229a92a31b5a9cf1d9b1162))
* **deps:** update module go.uber.org/dig to v1.19.0 ([#18](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/18)) ([53082ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/53082ff4b51d61dc15b38ab64c3959dd9bebb653))
* fix a test case name. ([b4f49c3](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b4f49c333135f60d5eda7b00006dc3e97e72c51d))
* install moq on release ([#13](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/13)) ([04da01c](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/04da01c77436b5c072f0bce964a90960e829254d))
* publish as repository owner ([#10](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/10)) ([45d71de](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/45d71de722cb11539d474a6050308638764fc204))
* release ([#29](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/29)) ([5c111e7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/5c111e714b25b9b8d08782d5313067ff6e82564a))
* renovate vonfig ([#14](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/14)) ([e7335ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/e7335ffa1e206102c15518826980770e29ec328a))
* run ci on release-please PRs ([#7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/7)) ([1974317](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/19743170ccf5fd7985a80d0a76177886ecaa789a))
* update GHA workflow to wait for all test jobs. ([6592646](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/65926463eb49c2b1100c893a1b5974fe3f29d9a9))

## [1.1.0](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/compare/gh-ratelimit-metrics-exporter-v1.0.0...gh-ratelimit-metrics-exporter-v1.1.0) (2026-05-07)


### Features

* add golangci-lint. ([9a057bd](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9a057bdb4a067d310f2d792185c8434d7f9369e0))
* add lint. ([61f21e4](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/61f21e45a0af555396b1e40a3f60eafe5be7c5c4))
* devcontainer wip ([b2b22ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b2b22ffc7839e95a9a7beb737a04fda65d6a5705))
* impl release flow and e2e test ([a750c6d](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/a750c6dd30dd33afc3c82e26cb441f587712e822))
* impl release flow and e2e test ([c4b54fb](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/c4b54fb46c40d6248e2599739a98a65f8ca6b7ca))
* run lint, test, e2e-test in parallel ([9ebc7f2](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9ebc7f22d1333f217842fcd62a72a4acb7f811ba))


### Bug Fixes

* **deps:** update module github.com/rs/zerolog to v1.35.1 ([#16](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/16)) ([00c3915](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/00c391578b42d0a91229a92a31b5a9cf1d9b1162))
* **deps:** update module go.uber.org/dig to v1.19.0 ([#18](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/18)) ([53082ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/53082ff4b51d61dc15b38ab64c3959dd9bebb653))
* fix a test case name. ([b4f49c3](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b4f49c333135f60d5eda7b00006dc3e97e72c51d))
* install moq on release ([#13](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/13)) ([04da01c](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/04da01c77436b5c072f0bce964a90960e829254d))
* publish as repository owner ([#10](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/10)) ([45d71de](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/45d71de722cb11539d474a6050308638764fc204))
* renovate vonfig ([#14](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/14)) ([e7335ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/e7335ffa1e206102c15518826980770e29ec328a))
* run ci on release-please PRs ([#7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/7)) ([1974317](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/19743170ccf5fd7985a80d0a76177886ecaa789a))
* update GHA workflow to wait for all test jobs. ([6592646](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/65926463eb49c2b1100c893a1b5974fe3f29d9a9))

## 1.0.0 (2026-05-07)


### Features

* add golangci-lint. ([9a057bd](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9a057bdb4a067d310f2d792185c8434d7f9369e0))
* add lint. ([61f21e4](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/61f21e45a0af555396b1e40a3f60eafe5be7c5c4))
* devcontainer wip ([b2b22ff](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b2b22ffc7839e95a9a7beb737a04fda65d6a5705))
* impl release flow and e2e test ([a750c6d](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/a750c6dd30dd33afc3c82e26cb441f587712e822))
* impl release flow and e2e test ([c4b54fb](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/c4b54fb46c40d6248e2599739a98a65f8ca6b7ca))
* run lint, test, e2e-test in parallel ([9ebc7f2](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/9ebc7f22d1333f217842fcd62a72a4acb7f811ba))


### Bug Fixes

* fix a test case name. ([b4f49c3](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/b4f49c333135f60d5eda7b00006dc3e97e72c51d))
* publish as repository owner ([#10](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/10)) ([45d71de](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/45d71de722cb11539d474a6050308638764fc204))
* run ci on release-please PRs ([#7](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/issues/7)) ([1974317](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/19743170ccf5fd7985a80d0a76177886ecaa789a))
* update GHA workflow to wait for all test jobs. ([6592646](https://github.com/air-hand-org/gh-ratelimit-metrics-exporter/commit/65926463eb49c2b1100c893a1b5974fe3f29d9a9))
