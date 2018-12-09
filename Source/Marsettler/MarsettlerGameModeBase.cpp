#include "MarsettlerGameModeBase.h"
#include "Networking.h"
#include "Sockets.h"
#include "message/actor_generated.h"

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
	flatbuffers::FlatBufferBuilder  builder;

	auto position = message::Vector(1.0f, 2.0f);
	auto actor = message::CreateActor(builder, &position);
	builder.Finish(actor);

	uint8_t *buf = builder.GetBufferPointer();

	auto readAcotr = message::GetActor(buf);
	auto readPosition = readAcotr->Position();

	UE_LOG(LogTemp, Log, TEXT("Hello flatbuffers! %f: %f"), readPosition->X(), readPosition->Y());
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
	}
}
