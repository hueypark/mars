#pragma once

#include "CoreMinimal.h"
#include "GameFramework/GameModeBase.h"
#include "MarsettlerGameModeBase.generated.h"

class FSocket;
struct FIPv4Address;

UCLASS()
class MARSETTLER_API AMarsettlerGameModeBase : public AGameModeBase
{
	GENERATED_BODY()

private:
	FSocket* socket = nullptr;

public:
	AMarsettlerGameModeBase();

	void BeginPlay() override;

	void Tick(float delta) override;

private:
	void connectSocket(const FIPv4Address& ipStr, const int32 port);
};
