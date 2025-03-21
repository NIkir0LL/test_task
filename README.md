# This task is a test task.

данное задание было взять c hh - https://vyksa.hh.ru/vacancy/118626902?query=%D0%9F%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B8%D1%81%D1%82+golang&hhtmFrom=vacancy_search_list

=== RUN   TestCreateUser

2025/03/21 12:04:25 CreateUser: Successfully created user - ID: 1, Name: Test User, Email: test@example.com

[GIN] 2025/03/21 - 12:04:25 | 201 |     582.776µs |                 | POST     "/users"

--- PASS: TestCreateUser (0.00s)

=== RUN   TestGetUser

2025/03/21 12:04:25 GetUser: Successfully retrieved user - ID: 1, Name: Test User, Email: test@example.com

[GIN] 2025/03/21 - 12:04:25 | 200 |      66.242µs |                 | GET      "/users/1"

--- PASS: TestGetUser (0.00s)

=== RUN   TestGetUserNotFound

2025/03/21 12:04:25 GetUser: User not found - ID: 999

[GIN] 2025/03/21 - 12:04:25 | 404 |      71.573µs |                 | GET      "/users/999"

--- PASS: TestGetUserNotFound (0.00s)

=== RUN   TestUpdateUser

2025/03/21 12:04:25 UpdateUser: Successfully updated user - ID: 1, Name: Updated User, Email: updated@example.com

[GIN] 2025/03/21 - 12:04:25 | 200 |     208.841µs |                 | PUT      "/users/1"

--- PASS: TestUpdateUser (0.00s)

=== RUN   TestUpdateUserInvalidID

2025/03/21 12:04:25 UpdateUser: Invalid ID - strconv.Atoi: parsing "invalid": invalid syntax

[GIN] 2025/03/21 - 12:04:25 | 400 |     119.677µs |                 | PUT      "/users/invalid"

--- PASS: TestUpdateUserInvalidID (0.00s)
