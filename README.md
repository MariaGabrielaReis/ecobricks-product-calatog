# EcoBricks [![license](https://img.shields.io/badge/LICENCE-MIT-blue)]() [![status](https://img.shields.io/badge/STATUS-DONE-brightgreen)]()

> Application developed in "Imersão Full Stack && Full Cycle" event, promoted by Full Cycle 🚀

**EcoBricks** is an e-commerce that sells different types of bricks, such as ecological ones! 🧱

<details>
  <summary>✨ <b>Features (of this repo)</b> </summary>

- [x] Create category
- [x] List categories
  - [x] All
  - [x] By ID
- [x] Create product
- [x] List products
  - [x] All
  - [x] By ID
  - [x] By category ID

</details>

> [!TIP]
> Visit [demo]() or the others repositories of this project:
>
> - **Frontend**: [@MariaGabrielaReis/ecobricks-web](https://github.com/MariaGabrielaReis/ecobricks-web) - `Next, Typescript, React Hook Form, Yup`
> - **API Checkout**: [@MariaGabrielaReis/ecobricks-checkout](https://github.com/MariaGabrielaReis/ecobricks-checkout) - `Nest, Typescript, TypeORM, MySQL, RabbitMQ, Docker`
> - **API Payment**: [@MariaGabrielaReis/ecobricks-payment](https://github.com/MariaGabrielaReis/ecobricks-payment) - `Go, RabbitMQ, Docker`

![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)

<div align="center">
  <img alt="EcoBricks" src="" />
</div>

## Building and running locally

Be sure you have [Git](https://git-scm.com/downloads), [Docker](https://www.docker.com/) and [Go](https://go.dev/) configurated, then follow the directions below.

```bash
# Clone this repository
$ git clone https://github.com/MariaGabrielaReis/ecobricks-product-catalog

# Access project folder
$ cd ecobricks-product-catalog
```

> [!IMPORTANT]
> Before running the project, set up the container with database:
>
> ```bash
> $ docker-compose up -d
> ```
>
> Now, access database with root user (the password is `root`):
>
> ```bash
> $ docker-compose exec mysql bash
> $ mysql -uroot -p
> ```
>
> Copy and paste the SQL instructions from `db.sql` file, then check if the product and category tables were created:
>
> ```bash
> $ show tables;
> ```

```bash
# Run
$ cd cmd/catalog && go run main.go
```

> [!NOTE]
> The server will start on port 8080 - <http://localhost:8080> <br>
>
> - For quick API testing, use `test.http` file to send requests with the VSCode "REST Client" extension.
