# CSWeekly 🖥️

A dynamic website for hosting weekly coding problems for SCSU Computer Science Club. Challenge yourself with new programming problems every week and improve your coding skills!

## 🚀 Features

- Weekly updated coding problems
- Clean, responsive UI powered by TailwindCSS and HyperUI
- Beatifully simple and lightweight Go backend using Echo framework
- Cloud-based SQLite database with Turso
- Deployed on Digital Ocean App Engine

## 💻 Tech Stack

- **Backend:** Go with Echo framework
- **Database:** Turso (Cloud SQLite)
- **Frontend:** TailwindCSS, HyperUI
- **Deployment:** Digital Ocean App Engine
- **Build:** Heroku/Go buildpack

## 📁 Project Structure

```
csweekly/
├── db/
│   └── db.go          # Database connection initialization
├── public/            # Static files directory
├── handlers.go        # DB-connected handler functions
└── main.go           # Server initialization and routing
```

## 🔧 Configuration

The following environment variables are required:

```
DATABASE_URL=your_turso_connection_string
PORT=8080
SECRETKEY = verysecret  # used for protecting the input form
```

The application is configured to deploy on Digital Ocean App Engine using the Heroku/Go buildpack. Follow these steps to deploy:


## 📝 License or whatever

Anyone can use, modify, and publish this website. Pull requests are welcome :)

## 🙏 Acknowledgements

- [Echo Framework](https://echo.labstack.com/)
- [TailwindCSS](https://tailwindcss.com/)
- [HyperUI](https://hyperui.dev/)
- [Turso](https://turso.tech/)
- SCSU Computer Science Club
