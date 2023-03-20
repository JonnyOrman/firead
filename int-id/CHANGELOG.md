# Changelog

## v0.1.0

- Read Firestore documents with either an `int` ID via a `/{id}` endpoint
    - Returns `200` response with document if found.
    - Returns `404` response with document if not found.