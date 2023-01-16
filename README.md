This is a microservice named Restful Library.

All the operations listed below has been implemented via REST API 
    
    1. Get to see all the books with avaiablity.
    2. Get to see one specific book is avalable to reserve or not.
    3. Discard one book from library if it got damaged.
    4. Reserve one book from library.
    5. Release one book to library.
    6. Search books by Author..

This microservice has been implemented in golang. All modules having their own uni test cases.

JenkinsFile has been written to make one pipeline with listed stages - 
    1. Build the code
    2. Test the code
    3. Lint the code
    4. Deploy