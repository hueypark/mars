#include "MarsettlerGameModeBase.h"
#include "Message/message_generated.h"

AMarsettlerGameModeBase::AMarsettlerGameModeBase()
{
	PrimaryActorTick.bStartWithTickEnabled = true;
	PrimaryActorTick.bCanEverTick = true;
}

void AMarsettlerGameModeBase::Tick(float DeltaTime)
{
	flatbuffers::FlatBufferBuilder  builder;

	auto position = Message::Vec(1.0f, 2.0f);
	auto actor = Message::CreateActor(builder, &position );
	builder.Finish(actor);

	uint8_t *buf = builder.GetBufferPointer();

	auto readAcotr = Message::GetActor(buf);
	auto readPosition = readAcotr->Position();

	UE_LOG(LogTemp, Warning, TEXT("Hello flatbuffers! %f: %f"), readPosition->X(), readPosition->Y() );
}
