# Migration from Direct PostgreSQL to API with Header Auth

## Summary
Successfully migrated the inspections reader application from direct PostgreSQL database access to using a REST API with UUID-based authentication.

## Changes Made

### 1. Environment Configuration (.env)
- **Added**: `API_URL=http://localhost:8081` - Base URL for the API
- **Added**: `READER_UUID=3b1eb00a-090e-46c1-9cf0-24c816bb0141` - UUID token for authentication
- **Kept**: `PDF_PATH` - Path to PDF files directory
- **Removed/Commented**: Database credentials (DB_USER, DB_PASS, DB_HOST, DB_NAME) - no longer needed

### 2. New API Client Module (parsers/api_client.go)
Created a new module to handle all API communication:
- **APIClient struct**: Manages HTTP client, base URL, and UUID token
- **NewAPIClient()**: Initializes client from environment variables
- **sendBatchRequest()**: Generic method to send POST requests with JSON payloads
- **SendInspectionsBatch()**: Sends inspection forms to `/inspections/batch`
- **SendDrySystemsBatch()**: Sends dry system forms to `/dry-systems/batch`
- **SendPumpSystemsBatch()**: Sends pump system forms to `/pump-systems/batch`
- **SendBackflowBatch()**: Sends backflow forms to `/backflow/batch`

Features:
- 60-second timeout for API requests
- Bearer token authentication in Authorization header
- JSON marshaling/unmarshaling
- Error handling with detailed messages
- Response parsing (inserted/updated/total counts)

### 3. Refactored main.go
**Removed**:
- Direct database connection (`parsers.ConnectDatabase()`)
- Database table creation calls (`Create*Table()`)
- Individual record insertion calls (`Insert*Table()`)
- Database connection cleanup (`db.Close()`)

**Added**:
- API client initialization
- Batch accumulator slices for each form type
- Mutex for thread-safe concurrent access to slices
- Batch sending after all PDFs are processed
- API response logging

**Flow Change**:
```
OLD: PDF → Parse → Insert to DB (one at a time)
NEW: PDF → Parse → Accumulate in memory → Send batches to API
```

### 4. Preserved Functionality
- PDF file discovery and processing
- Worker pool with 10 concurrent workers
- Progress bar display
- Form type detection and mapping
- Backflow choice processing
- Logging to inspections.log

## API Endpoints Used

All endpoints require `Authorization: Bearer <UUID>` header.

### POST /inspections/batch
Accepts array of complete InspectionForm objects (200+ fields)

### POST /dry-systems/batch
Accepts array of complete DryForm objects

### POST /pump-systems/batch
Accepts array of complete PumpForm objects

### POST /backflow/batch
Accepts array of complete BackflowForm objects

## API Response Format
```json
{
  "inserted": 10,
  "updated": 5,
  "total": 15
}
```

## Benefits of Migration

1. **Centralized Authentication**: UUID-based auth managed by API
2. **Better Separation of Concerns**: Reader app focuses on PDF parsing, API handles data persistence
3. **Improved Performance**: Batch operations reduce network overhead
4. **Easier Maintenance**: Database schema changes only affect API
5. **Better Error Handling**: API provides structured error responses
6. **Scalability**: Multiple reader instances can use same API

## Configuration

### Before Running
1. Ensure API server is running on the configured URL
2. Verify UUID token is valid and not revoked in the users table
3. Update `API_URL` and `READER_UUID` in .env file with production values

### Testing
```bash
# Build the application
go build -o bin\inspections_reader.exe

# Run the application
.\bin\inspections_reader.exe
```

## Troubleshooting

### "API_URL not set in environment"
- Check .env file exists in the application directory
- Verify API_URL is set correctly

### "READER_UUID not set in environment"
- Check .env file exists
- Verify READER_UUID is set correctly

### "API returned status 401"
- UUID token is invalid or revoked
- Check users table in database
- Verify Authorization header format

### "API returned status 500"
- API server error
- Check API server logs
- Verify database connection on API side

### Connection timeout
- API server may be down
- Check API_URL is correct
- Verify network connectivity
- Consider increasing timeout in api_client.go

## Rollback Plan

If needed to rollback to direct database access:
1. Uncomment database credentials in .env
2. Restore original main.go from git history
3. Rebuild application

## Future Enhancements

- Add retry logic with exponential backoff
- Implement chunking for very large batches
- Add progress reporting during API uploads
- Support for partial batch failures
- Health check before processing PDFs
