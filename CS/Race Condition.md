___
## 정의

- 공유 자원에 대해 여러 프로세스가 동시에 접근을 시도할 때, 타이밍이나 순서 등이 결과값에 영향을 줄 수 있는 상태

---
## Deadlock

- 여러 프로세스나 스레드가 절대 일어나지 않는 이벤트나 자원 할당을 위해 무한정 대기하는 상태
- 데드락이 발생하면 프로세스는 blocked 상태에서 무한히 대기한다

---
## Starvation

- 프로세스가 CPU 자원의 할당을 무한히 대기하는 상태
- CPU 스케줄링과 밀접한 관련이 있다

---
## Critical Section(임계영역)

- 정의
  - 운영체제에서 여러 프로세스가 데이터를 공유하면서 수행될 때 각 프로세스에서 공유 자원에 접근하는 프로그램 코드 부분을 의미한다
- 요구조건
  - Mutual Exclusion (상호 배제)
    - 한 프로세스가 자신의 critical section이면 다른 프로세스들은 critical section에 진입할 수 없다
  - Progress (진행)
    - critical section이 비어있으면 자원을 사용할 수 있어야 한다(deadlock free)
  - Bounded Waiting (유한 대기)
    - 프로세스는 유한 시간 내에 critical section에 진입할 수 있어야 한다(starvation free)

---
## Critical Section Problem 해결

### Peterson's Algorithm

### Mutex

- 공유된 자원의 데이터를 여러 스레드가 접근하는 것을 막는 방법
- Critical Section을 가진 쓰레드들의 Running time이 서로 겹치지 않게 각각 단독으로 실행되게 하는 기술
- 동기화 대상이 오로지 1개일 때 사용
- 자원을 소유할 수 있고 책임을 가진다
- 뮤텍스는 Locking 메커니즘으로 락을 건 스레드만이 임계 영역을 나갈때 락을 해제할 수 있다

### Semaphore

- 공유된 자원의 데이터를 여러 프로세스가 접근하는 것을 막는 방법
- 리소스의 상태를 나타내는 간단한 카운터
- 일반적으로 비교적 긴 시간을 확보하는 리소스에 대해 사용
- 하나의 스레드/프로세스만 들어가게 할 수도 있고 여러 개의 스레드/프로세스가 들어가게 할 수도 있다(동기화 대상이 1개 이상일 때 사용)
- 자원을 소유할 수 없다
- 세마포어는 Signaling 메커니즘으로 락을 걸지 않은 스레드도 signal을 사용해 락을 해제할 수 있다

---
## References

- https://velog.io/@klloo/%EC%9A%B4%EC%98%81%EC%B2%B4%EC%A0%9C-%EA%B2%BD%EC%9F%81-%EC%83%81%ED%83%9C-Race-Condition
- https://charles098.tistory.com/88
- https://iredays.tistory.com/125
- https://itandhumanities.tistory.com/66
- https://heeonii.tistory.com/14
