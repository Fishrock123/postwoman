package handlers

import (
    "context"
    "encoding/json"

    "github.com/labstack/echo/v4"
    "github.com/stephenafamo/scan"
    "github.com/stephenafamo/scan/pgxscan"

    "postwoman/models"
)

func GetAllRequests(c echo.Context) error {

    email := c.Param("email")

    res, err := pgxscan.All(context.Background(), db, scan.StructMapper[models.Request](), "SELECT * FROM request WHERE user_email = $1 AND hidden = false", email)

    if err != nil {
        return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
    }

    if len(res) == 0 {
        return c.JSONPretty(404, errorJSON("User Error", "No requests found from this user email"), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func CreateRequest(c echo.Context) error {

    var res models.Request
    var data models.Request
    email := c.Param("email")

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        err := db.QueryRow(context.Background(), "INSERT INTO request (user_email, url, method, origin, headers, body, status, date, hidden) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *",
            email, data.Url, data.Method, data.Origin, data.Headers, data.Body, data.Status, data.Date, data.Hidden).Scan(&res.ID, &res.User_email, &res.Url, &res.Method, &res.Origin, &res.Headers, &res.Body, &res.Status, &res.Date, &res.Hidden)

        if err != nil {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }

        if res.ID == 0 {
            return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func HideRequest(c echo.Context, email string) map[string]interface{} {

    var res string
    requestID := c.Param("reqID")

    err := db.QueryRow(context.Background(), "UPDATE request SET hidden = true WHERE id = $1 AND user_email = $2 RETURNING id", requestID, email).Scan(&res)

    if res != requestID {
        return map[string]interface{}{"status": 401, "res": errorJSON("User Error", "No requests found made with the email and id provided")}
    }

    if err != nil {
        return map[string]interface{}{"status": 401, "res": errorJSON("Server Error", err.Error())}
    }

    return map[string]interface{}{"status": 401, "res": "successful"}
}

func DeleteRequest(c echo.Context) error {

    var res string
    requestID := c.Param("reqID")
    email := c.Param("email")

    err := db.QueryRow(context.Background(), "DELETE FROM request WHERE id = $1 AND user_email = $2 RETURNING id", requestID, email).Scan(&res)

    if res != requestID {
        return c.JSONPretty(404, errorJSON("User Error", "No requests found made with the email and id provided"), " ")
    }

    if err != nil {
        return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
    }

    return c.NoContent(200)
}
