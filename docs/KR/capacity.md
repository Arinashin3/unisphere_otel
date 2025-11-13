# Capacity Collector
DELL EMC Unity System의 용량 정보를 수집합니다.  
`metric`, `log`, `trace` 중 `metric`으로만 나타납니다.

## Metric Info
| Metric 명                                 | 값 | 유닛 | 추가 라벨 | 설명                           |
|------------------------------------------|---|----|-------|------------------------------|
| unisphere_capacity_total_capacity        | - | mb | -     | 시스템의 총 용량                    |
| unisphere_capacity_used_capacity         | - | mb | -     | 시스템에서 현재 사용중인 용량             |
| unisphere_capacity_free_capacity         | - | mb | -     | 시스템에서 현재 잔여 용량               |
| unisphere_capacity_preallocated_capacity | - | mb | -     | 시스템에서 현재 사전 할당된 용량           |
| unisphere_capacity_total_provision       | - | mb | -     | 시스템에서 스냅샷을 포함한 Provision된 용량 |

## Sample) Alert Rule

### Percentage Serverity Example
 - warning (80%)
 - critical (90%)
 - fatal (100%)
```yaml
groups:
  - name: UnisphereHighCapacity
    labels:
      component: 'capacity'
    rules:
      # High Capacity = Warning(80%)
      - alert: 'UnisphereHighCapacity'
        expr: '(unisphere_capacity_used_capacity_mb / unisphere_capacity_total_capacity_mb) * 100 >= 80'
        labels:
          severity: 'warning'
        annotations:
          summary: '{{ $labels.host_name }} capacity is too high. (current={{ $value.humanize }}%)'
          
      # High Capacity = Critical(90%)
      - alert: 'UnisphereHighCapacity'
        expr: '(unisphere_capacity_used_capacity_mb / unisphere_capacity_total_capacity_mb) * 100 >= 90'
        labels:
          severity: 'critical'
        annotations:
          summary: '{{ $labels.host_name }} capacity is too high. (current={{ $value.humanize }}%)'
          
      # High Capacity = Fatal(100%)
      - alert: 'UnisphereHighCapacity'
        expr: '(unisphere_capacity_used_capacity_mb / unisphere_capacity_total_capacity_mb) * 100 >= 100'
        labels:
          severity: 'critical'
        annotations:
          summary: '{{ $labels.host_name }} capacity is too high. (current={{ $value.humanize }}%)'
```

### Calculate Percentage and return Free Capacity Example
- warning (80%)
- critical (90%)
- fatal (100%)
```yaml
groups:
  - name: UnisphereHighCapacity
    labels:
      component: 'capacity'
    rules:
      # High Capacity = Warning(80%)
      - alert: 'UnisphereHighCapacityOutputFreeMB'
        expr: '((unisphere_capacity_used_capacity_mb / unisphere_capacity_total_capacity_mb) * 100 >= 80) * 0 + unisphere_capacity_free_capacity_mb'
        labels:
          severity: 'warning'
        annotations:
          summary: '{{ $labels.host_name }} capacity is too high. (current={{ $value.humanize1024 }} MB)'
          
      # High Capacity = Critical(90%)
      - alert: 'UnisphereHighCapacityOutputFreeMB'
        expr: '((unisphere_capacity_used_capacity_mb / unisphere_capacity_total_capacity_mb) * 100 >= 90) * 0 + unisphere_capacity_free_capacity_mb'
        labels:
          severity: 'critical'
        annotations:
          summary: '{{ $labels.host_name }} capacity is too high. (current={{ $value.humanize1024 }} MB)'
          
      # High Capacity = Fatal(100%)
      - alert: 'UnisphereHighCapacityOutputFreeMB'
        expr: '((unisphere_capacity_used_capacity_mb / unisphere_capacity_total_capacity_mb) * 100 >= 100) * 0 + unisphere_capacity_free_capacity_mb'
        labels:
          severity: 'critical'
        annotations:
          summary: '{{ $labels.host_name }} capacity is too high. (current={{ $value.humanize1024 }} MB)'
```

## Sample) Dashboard
