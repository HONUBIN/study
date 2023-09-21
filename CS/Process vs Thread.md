- - -
## Process
- 정의
	- 메모리에 적재되어 실행되고 있는 프로그램의 인스턴스
- 특징
	- 독립된 메모리 영역(Code, Data, Stack, Heap의 구조)을 할당받는다
	- 프로세스당 최소 1개의 스레드(메인 스레드)를 가진다
	- 다른 프로세스의 자원에 접근하려면 프로세스 간의 통신(IPC, inter-process-communication)을 사용
	  EX. 파이프, 파일, 소켓
- - -
## Thread
- 정의
	- 프로세스 내에서 실행되는 흐름의 단위
- 특징
	- 프로세스 내에서 각각 Stack만 따로 할당받고 Code, Data, Heap 영역은 공유한다
- - -
## 멀티 프로세스와 멀티 스레드의 차이
### 멀티 프로세스
- 장점
	- 여러 개의 자식 프로세스 중 하나에 문제가 발생하더라도 다른 프로세스에 영향을 주지 않는다
- 단점
	- Context Switching 오버헤드
	- 복잡한 통신(IPC)
### 멀티 스레드
- 장점
	- 자원 소모 감소(Code, Data, Heap을 공유하여 메모리 공간과 시스템 콜이 줄어듦)
	- Context Switching 비용 감소/속도 증가
	- Stack을 제외한 모든 메모리 영역을 공유하여 통신 비용 감소/속도 증가
- 단점
	- 자원을 공유하기 때문에 동기화 문제가 발생할 수 있다
	- 하나의 스레드에 문제가 발생하면 같은 프로세스 내의 다른 스레드가 영향을 받는다
- - -
## References
- https://gmlwjd9405.github.io/2018/09/14/process-vs-thread.html
- https://inpa.tistory.com/entry/%F0%9F%91%A9%E2%80%8D%F0%9F%92%BB-%ED%94%84%EB%A1%9C%EC%84%B8%EC%8A%A4-%E2%9A%94%EF%B8%8F-%EC%93%B0%EB%A0%88%EB%93%9C-%EC%B0%A8%EC%9D%B4#%ED%94%84%EB%A1%9C%EC%84%B8%EC%8A%A4_process