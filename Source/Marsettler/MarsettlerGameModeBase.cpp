#include "MarsettlerGameModeBase.h"

AMarsettlerGameModeBase::AMarsettlerGameModeBase()
{
	PrimaryActorTick.bStartWithTickEnabled = true;
	PrimaryActorTick.bCanEverTick = true;
}

void AMarsettlerGameModeBase::Tick(float DeltaTime)
{
	UE_LOG(LogTemp, Warning, TEXT("Hello, World!"));
}
