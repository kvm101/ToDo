# ToDo API
**version 5.0.0-alfa (10.01.25)**

## 1.Desciption
1.1.Technologies to used:
  - **Golang**
  - **Postgres**
  - **Python**
  - **Streamlit**

## 2.Project version:
  version 1.0.0-alfa (27.12.24):
  - Marked project structure.
  - Connected DB.

  version 1.1.0-alfa (27.12.24):
  - Created db functions (Add, Delete, Edit, Read).
  - Decomposition api from main.go.

  version 1.1.1-alfa (29.12.24):
  - Upgrade db functions (Add, Delete, Edit, Read).

  version 1.2.0-alfa (30.12.24):
  - JSON transfer.
  - Edit project interface instead of ConsoleApp will be WebApp.
  - added technology: React (JavaScript).

  version 1.2.1-alfa (30.12.24):
  - Fixed function read() through json incorrection.

  version 1.3.0-alfa (30.12.24):
  - Added Logging.

  version 1.4.0-alfa (30.12.24):
  - Added Sections(Tasks, Done tasks).

  version 1.5.0-alfa (01.01.25):
  - Added Sections (Date, Complexity, Importance).
  - Added Sorting for date, complexity, importance in function readTasks for 2 parameters (section - "all", "undone", "done";
  sortf = "date_asc", "date_desc", "head_asc", "head_desc", "complexity_asc", "complexity_desc", "importance_asc", "importance_desc").

  version 1.5.1-alfa (02.01.25):
  - Fixed incorrect types of data.

  version 1.5.2-alfa (02.01.25):
  - Deleted sorting for switch logic.
  - Added sorting for parameter that enter manually.

  version 2.0.0-alfa (04.01.25):
  - Added Basic HTTP Authorization.

  version 2.0.1-alfa (04.01.25):
  - Fixed error of authentification and queries.

  version 2.1.0-alfa (04.01.25):
  - Added user registration.

  version 2.1.1-alfa (05.01.25):
  - Fixed bug that you could create any number of equal accounts

  version 3.0.0-alfa (05.01.25):
  - Added validator input data of registration

  version 4.0.0-alfa (06.01.25):
  - Added URL query parameters "limit" and "page"
  - Essues: have a bug with incorrect output

  version 4.0.1-alfa (06.01.25):
  - Fixed incorrect data output in readTasks()

  version 5.0.0-alfa (10.01.25):
  - Added general structure telemetry admin pannel on Streamlit

  planning in the next version(6.0.0)-alfa (20.01.25):
  - Replace "log" to "go.uber.org/zap"
  - Union "zap" with Python App Telemetry on Streamlit

  planning in version(7.0.0)-alfa:
  - Replace "github.com/lib/pq" to "github.com/jackc/pgx/v5"

  planning in version(8.0.0)-alfa:
  -Replace "net/http" to "github.com/gofiber/fiber"
  
  planning in future versions:
  - replaced http basic authorization to JWT.
  - Add OAuth 2.0.


## License
This project have [MIT License](LICENSE).
