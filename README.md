# Tfinit

<div>
<a href="https://codeclimate.com/github/paloth/tfinit/maintainability"><img src="https://api.codeclimate.com/v1/badges/8be4099c6dd4d6459841/maintainability" /></a>
<a href="https://codeclimate.com/github/paloth/tfinit/test_coverage"><img src="https://api.codeclimate.com/v1/badges/8be4099c6dd4d6459841/test_coverage" /></a>
</div>

A tools to inititate Terraform project

## Commands

tfinit:
  - [ ] module
    - [ ] local
      - [ ] name
      - [ ] provider
      - [ ] version
    - [ ] no arg
      - [ ] name
      - [ ] provider
      - [ ] version
  - [ ] projet
    - [ ] name
    - [ ] provider
    - [ ] version

## Files to create

### projet

```
projet
├── module
│   └── <provider>-<module_name>
│       ├── variables.tf
│       ├── version.tf
│       ├── outputs.tf
│       ├── main.tf
│       ├── README.md
│       └── datasources.tf
├── .gitignore
├── terraform
│   ├── variables.tf
│   ├── providers.tf
│   ├── outputs.tf
│   ├── main.tf
│   └── datasources.tf
├── LICENSE
└── README.md
```

### module

#### local

```
module
└── <provider>-<module_name>
    ├── variables.tf
    ├── version.tf
    ├── outputs.tf
    ├── main.tf
    ├── README.md
    └── datasources.tf
```

#### no arg

```
terraform-<provider>-<module_name>
├── variables.tf
├── version.tf
├── outputs.tf
├── main.tf
├── README.md
├── datasources.tf
├── LICENSE
└── .gitignore
```
