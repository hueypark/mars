#pragma once

#include "CoreMinimal.h"
#include "GameFramework/GameModeBase.h"
#include "MarsettlerGameModeBase.generated.h"

UCLASS()
class MARSETTLER_API AMarsettlerGameModeBase : public AGameModeBase
{
	GENERATED_BODY()
	
	AMarsettlerGameModeBase();

	void Tick(float DeltaTime);
};
