# GRAIL interview

This is my solution for the GRAIL interview question. It's admittedly a bit over the top for what should have been a short project, but I took the opportunity to
learn stuff including:

- The first Go code I've ever written
- The first time I've used Bazel outside Google's monolithic repository, which has some non-trivial differences
- The first time I've created a gRPC server not using Google internal-only frameworks

Other than that, I've sort-of used gRPC before with Java servers, but Google has a bunch of internal-only frameworks to build them (which is why I decided to learn
something new, since there's no RPC framework external to Google I've used before).

The API also follows best practices from https://aip.dev.

## Running instructions

Requires [Bazel](https://docs.bazel.build/versions/4.0.0/getting-started.html) to build. Run the server using the following command:

```shell
$ bazel run //cmd/participant-server -- -logtostderr
```

This starts a REST server listening on port `50051`. Note that everything is stored in memory, there is no persistance, so restarting the server will clear
everything.

## REST API

### Participant definition

Participants are represented as JSON objects with the following fields:

| **Field**   | **Description**                                                                                                                   |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------|
| name        | The **resource name** of the Participant, not the name of the person. See https://aip.dev/122#fields-representing-resource-names. |
| givenName   | The participant's given name.                                                                                                     |
| familyName  | The participant's family name.                                                                                                    |
| birthDate   | The participant's date of birth represented as an object with a `year`, `month` and `day`.                                        |
| phoneNumber | The participant's phone number.                                                                                                   |
| address     | The participant's physical address.                                                                                               |

### POST /v1/participants

Creates a new participant. The body of the request can include a JSON participant but must not set the `name` field as that will be generated. Returns JSON for
the created participant. Not all fields have to be set in case the information is collected gradually.

Example:

```shell
$ curl -X POST localhost:50051/v1/participants -d '{"givenName": "Graham", "familyName": "Rogers"}'
```

### GET /v1/participants/{reference}

Gets a participant. Should use the `name` of a previously created participant otherwise it will 404. Returns a JSON participant.

Example:

```shell
$ curl -X GET localhost:50051/v1/participants/AAA-001
```

### PATCH /v1/participants/{reference}

Updates a participant's details. The body of the request should be a JSON participant without the name field. Returns the JSON of the updated participant.

Example:

```shell
curl -X PATCH localhost:50051/v1/participants/AAA-001 -d '{"phoneNumber": "01189998819991197253"}'
```

### DELETE /v1/participants/{reference}

Deletes the participant, returning an empty JSON object.

Example:

```shell
curl -X DELETE localhost:50051/v1/participants/AAA-001
```
