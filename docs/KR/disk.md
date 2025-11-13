# Disk Collector
Disk의 상태와 용량을 수집합니다.  
`metric`, `log`, `trace` 중 `metric`으로만 나타납니다.

## Metric Info
| Metric 명                 | 값                                                                                                                        | 유닛  | 추가 라벨                                        | 설명                             |
|--------------------------|--------------------------------------------------------------------------------------------------------------------------|-----|----------------------------------------------|--------------------------------|
| unisphere_disk_health    | 0: Unknown<br/>5: OK<br/>7: OK_BUT<br/>10: DEGRADED<br/>15: MINOR<br/>20: MAJOR<br/>25: CRITICAL<br/>30: NON_RECOVERABLE | -   | `disk_id` `slot_id`                          | Disk의 현재 상태                    |
| unisphere_disk_info      | 1: emcPartNumber 감지<br/>0: emcPartNumber 미감지                                                                             | -   | `disk_id` `slot_id` `disk_model` `disk_part` | Disk의 모델 및 파츠 정보               |
| unisphere_disk_size      | -                                                                                                                        | mb | `disk_id` `slot_id`                          | Disk의 사이즈                      |
| unisphere_disk_is_in_use | 1: 사용중<br/>0: 미사용                                                                                                        | - | `disk_id` `slot_id`                          | Disk가 유저에 의해 쓰여진 데이터가 존재하는지 여부 |

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
  - name: UnisphereDiskHealthStatus
    labels:
      component: 'health'
    rules:
      # Disk Status == OK_BUT
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 7'
        labels:
          severity: 'OK_BUT'
        annotations:
          summary: '{{ $labels.host_name }} disk is OK_BUT(warning). (slot={{ $labels.slot_id }})'
          
      # Disk Status == DEGRADED
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 10'
        labels:
          severity: 'DEGRADED'
        annotations:
          summary: '{{ $labels.host_name }} disk is DEGRADED. (slot={{ $labels.slot_id }})'
          
      # Disk Status == MINOR
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 15'
        labels:
          severity: 'MINOR'
        annotations:
          summary: '{{ $labels.host_name }} disk is MINOR. (slot={{ $labels.slot_id }})'
          
      # Disk Status == MAJOR
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 20'
        labels:
          severity: 'MAJOR'
        annotations:
          summary: '{{ $labels.host_name }} disk is MAJOR. (slot={{ $labels.slot_id }})'
          
      # Disk Status == CRITICAL
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 25'
        labels:
          severity: 'CRITICAL'
        annotations:
          summary: '{{ $labels.host_name }} disk is CRITICAL. (slot={{ $labels.slot_id }})'
          
      # Disk Status == NON-RECOVERABLE
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 30'
        labels:
          severity: 'NON-RECOVERABLE'
        annotations:
          summary: '{{ $labels.host_name }} disk is NON-RECOVERABLE. (slot={{ $labels.slot_id }})'
```

#### Custom Serverity Example
 - warning
 - minor
 - critical
 - fatal
```yaml
groups:
  - name: UnisphereDiskHealthStatus
    labels:
      component: 'health'
    rules:
      # Disk Status == OK_BUT
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 7'
        labels:
          severity: 'warning'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} disk is warning. (slot={{ $labels.slot_id }})'

      # Disk Status == DEGRADED, MINOR
      - alert: 'UnisphereDiskHealthStatus'
        expr: '10 <= unisphere_disk_health <= 15'
        labels:
          severity: 'minor'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} disk is degraded. (slot={{ $labels.slot_id }})'

      # Disk Status == MAJOR, CRITICAL
      - alert: 'UnisphereDiskHealthStatus'
        expr: '20 <= unisphere_disk_health <= 25'
        labels:
          severity: 'critical'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} disk is critical. (slot={{ $labels.slot_id }})'

      # Disk Status == NON-RECOVERABLE
      - alert: 'UnisphereDiskHealthStatus'
        expr: 'unisphere_disk_health == 30'
        labels:
          severity: 'fatal'
        annotations:
          summary: 'DELL Unity: {{ $labels.host_name }} disk is fatal. (slot={{ $labels.slot_id }})'
```

## Sample) Dashboard