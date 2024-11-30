there are 3 operation paradigms:

1. command line
2. web page
3. REST API

the web page shares an in-memory database with the REST API, but the command line does not.

example REST API calls:

```
# create a new task
curl -X PUT http://172.22.218.64:5000/task \
  --data '{
  "Action": "Add",
  "Number": 0,
  "Item": "wash dishes",
  "Status": "Not Started"
}'

# read existing task
curl http://172.22.218.64:5000/task/1
```