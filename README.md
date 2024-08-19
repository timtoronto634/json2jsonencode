# json2jsonencode

## QuickStart 

```
go build .
cat sample.json | ./jsonencode
```

## Usage

parsing the AWS CloudWatch Dashboard Query to before-jsonencode string

```
aws cloudwatch get-dashboard --dashboard-name dashboard-name --query 'DashboardBody' --output json | json2jsonencode
```