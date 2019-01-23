#include "MarsettlerGameModeBase.h"
#include "Networking.h"
#include "Sockets.h"
#include "message/actor_generated.h"
#include "message/constants.h"
#include "message/login_generated.h"

AMarsettlerGameModeBase::AMarsettlerGameModeBase()
{
	PrimaryActorTick.bStartWithTickEnabled = true;
	PrimaryActorTick.bCanEverTick = true;
}

void AMarsettlerGameModeBase::BeginPlay()
{
	connectSocket(FIPv4Address(127, 0, 0, 1), 8080);
}

void AMarsettlerGameModeBase::Tick(float delta)
{
	TArray<uint8> head;
	head.Init(0, 8);
	int32 read;
	socket->Recv(head.GetData(), head.Num(), read);

	int32 id, size;
	FMemory::Memcpy(&id,   &head.GetData()[0], 4);
	FMemory::Memcpy(&size, &head.GetData()[4], 4);

	UE_LOG(LogTemp, Log, TEXT("Hello tcp! {id: %d, size:%d}"), id, size);

	TArray<uint8> body;
	body.Init(0, size);
	socket->Recv(body.GetData(), body.Num(), read);

	const fbs::Actor* actor = fbs::GetActor(body.GetData());
	UE_LOG(LogTemp, Log, TEXT("Hello tcp! {id: %lld, x: %f, y: %f}"), actor->Id(), actor->Position()->X(), actor->Position()->Y());
}

void AMarsettlerGameModeBase::connectSocket(const FIPv4Address& ip, const int32 port)
{
	socket = ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->CreateSocket(NAME_Stream, TEXT("default"), false);

	TSharedRef<FInternetAddr> addr = ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->CreateInternetAddr();
	addr->SetIp(ip.Value);
	addr->SetPort(port);

	if (socket->Connect(*addr))
	{
		UE_LOG(LogTemp, Log, TEXT("Socket connected."));

		const MessageID id   = MessageID::Login;
		const int32     size = sizeof(fbs::Login);
		TArray<uint8> head;
		head.Init(0, HeadSize);
		FMemory::Memcpy(&head.GetData()[0], &id,   4);
		FMemory::Memcpy(&head.GetData()[4], &size, 4);

		int32 sent;
		socket->Send(head.GetData(), HeadSize, sent);

		fbs::Login login(12321);
		socket->Send((const uint8*)(&login), size, sent);
	}
}
