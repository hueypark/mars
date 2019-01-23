#pragma once

#include "CoreMinimal.h"

const int HeadSize = 8;

enum class MessageID : uint8
{
	Actor = 0,
	Login,
	LoginResult,
};
