# DPE Collector
Disk Processor Enclosures의 상태와 최근 온도를 수집합니다.  
`metric`, `log`, `trace` 중 `metric`으로만 나타납니다.

## Metric Info
| Metric 명                          | 값                                                                                                                        | 유닛 | 추가 라벨    | 설명         |
|-----------------------------------|--------------------------------------------------------------------------------------------------------------------------|---|----------|------------|
| unisphere_dpe_health              | 0: Unknown<br/>5: OK<br/>7: OK_BUT<br/>10: DEGRADED<br/>15: MINOR<br/>20: MAJOR<br/>25: CRITICAL<br/>30: NON_RECOVERABLE | - | `dpe_id` | DPE의 현재 상태 |
| unisphere_dpe_current_temperature | -                                                                                                                        | - | `dpe_id` | DPE의 현재 온도 |

## Sample) Alert Rule

#### Original Serverity Example
- OK_BUT(warning)
- DEGRADED
- MINOR
- MAJOR
- CRITICAL
- NON-RECOVERABLE
```yaml
groups:
  - name: UnisphereDPEHealthStatus
    labels:
      component: 'health'
    rules:
      # DPE Status == OK_BUT
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 7'
        labels:
          severity: 'OK_BUT'
        annotations:
          summary: '{{ $labels.host_name }} dpe is OK_BUT(warning). (id={{ $labels.dpe_id }})'
          
      # DPE Status == DEGRADED
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 10'
        labels:
          severity: 'DEGRADED'
        annotations:
          summary: '{{ $labels.host_name }} dpe is DEGRADED. (id={{ $labels.dpe_id }})'
          
      # DPE Status == MINOR
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 15'
        labels:
          severity: 'MINOR'
        annotations:
          summary: '{{ $labels.host_name }} dpe is MINOR. (id={{ $labels.dpe_id }})'
          
      # DPE Status == MAJOR
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 20'
        labels:
          severity: 'MAJOR'
        annotations:
          summary: '{{ $labels.host_name }} dpe is MAJOR. (id={{ $labels.dpe_id }})'
          
      # DPE Status == CRITICAL
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 25'
        labels:
          severity: 'CRITICAL'
        annotations:
          summary: '{{ $labels.host_name }} dpe is CRITICAL. (id={{ $labels.dpe_id }})'
          
      # DPE Status == NON-RECOVERABLE
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 30'
        labels:
          severity: 'NON-RECOVERABLE'
        annotations:
          summary: '{{ $labels.host_name }} dpe is NON-RECOVERABLE. (id={{ $labels.dpe_id }})'
```

#### Custom Serverity Example
- warning
- minor
- critical
- fatal
```yaml
groups:
  - name: UnisphereDPEHealthStatus
    labels:
      component: 'health'
    rules:
      # DPE Status == OK_BUT
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 7'
        labels:
          severity: 'warning'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} dpe is warning. (id={{ $labels.dpe_id }})'

      # DPE Status == DEGRADED, MINOR
      - alert: 'UnisphereDPEHealthStatus'
        expr: '10 <= unisphere_dpe_health <= 15'
        labels:
          severity: 'minor'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} dpe is degraded. (id={{ $labels.dpe_id }})'

      # DPE Status == MAJOR, CRITICAL
      - alert: 'UnisphereDPEHealthStatus'
        expr: '20 <= unisphere_dpe_health <= 25'
        labels:
          severity: 'critical'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} dpe is critical. (id={{ $labels.dpe_id }})'

      # DPE Status == NON-RECOVERABLE
      - alert: 'UnisphereDPEHealthStatus'
        expr: 'unisphere_dpe_health == 30'
        labels:
          severity: 'fatal'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} dpe is fatal. (id={{ $labels.dpe_id }})'
```

## Sample) Dashboard
