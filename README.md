# json2jsonencode
Convert json string to before-jsonencode string exclusively for converting AWS CloudWatch Dashboard's DashboardBody (json string) to jsonencode({}) style string
The output can be pasted as terraform code of "dashboard_body"

## Installation

```
go install github.com/timtoronto634/json2jsonencode
```

## Usage

```
aws cloudwatch get-dashboard --dashboard-name dashboard-name --query 'DashboardBody' --output json | json2jsonencode
```


## Sample Input/Output

input
```
{
  "widgets": [
    {
      "height": 1,
      "width": 16,
      "y": 0,
      "x": 0,
      "type": "text",
      "properties": {
        "markdown": "# api\n"
      }
    },
    {
      "height": 6,
      "width": 6,
      "y": 1,
      "x": 0,
      "type": "metric",
      "properties": {
        "annotations": {
          "alarms": [
            "arn:aws:cloudwatch:ap-northeast-1:511111:alarm:api-cpu-utilization-rate"
          ]
        },
        "stacked": false,
        "title": "api-cpu-utilization-rate",
        "type": "chart",
        "view": "timeSeries"
      }
    }
  ]
}
```

output

```
jsonencode({
  widgets = [
    {
      properties = {
        markdown = <<-MARKDOWN
# api
MARKDOWN
      }
      height = 1
      width = 16
      y = 0
      x = 0
      type = "text"
    },
    {
      height = 6
      width = 6
      y = 1
      x = 0
      type = "metric"
      properties = {
        annotations = {
          alarms = [
            "arn:aws:cloudwatch:ap-northeast-1:511111:alarm:api-cpu-utilization-rate",
        ]
        }
        stacked = false
        title = "api-cpu-utilization-rate"
        type = "chart"
        view = "timeSeries"
      }
    },
]
})
```
