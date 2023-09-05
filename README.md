# worldupgrader

Worldupgrader upgrades Minecraft Bedrock Edition worlds and their components to the latest version.

## Documentation

[![PkgGoDev](https://pkg.go.dev/badge/github.com/df-mc/worldupgrader)](https://pkg.go.dev/github.com/df-mc/worldupgrader)

## Credits

Block state and item upgrading is done by using PMMP's generated upgrade schemas. These schemas are stored in this
repository using subtrees, which can be updated using the following commands. If you want to target a specific branch,
then replace master with the name of the branch you wish to use.

```shell
git subtree pull --prefix blockupgrader/remote https://github.com/pmmp/BedrockBlockUpgradeSchema.git master --squash
```
```shell
git subtree pull --prefix itemupgrader/remote https://github.com/pmmp/BedrockItemUpgradeSchema.git master --squash
```