using UnrealBuildTool;
using System.Collections.Generic;

public class MarsettlerTarget : TargetRules
{
	public MarsettlerTarget(TargetInfo Target) : base(Target)
	{
		Type = TargetType.Game;

		ExtraModuleNames.AddRange( new string[] { "Marsettler" } );
	}
}
