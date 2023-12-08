rm -rf blockupgrader/schemas itemupgrader/schemas
git clone https://github.com/pmmp/BedrockBlockUpgradeSchema
mv BedrockBlockUpgradeSchema/nbt_upgrade_schema blockupgrader/schemas
git clone https://github.com/pmmp/BedrockItemUpgradeSchema
mv BedrockItemUpgradeSchema/id_meta_upgrade_schema itemupgrader/schemas

BLOCK_COMMIT=$(git -C BedrockBlockUpgradeSchema rev-parse HEAD)
ITEM_COMMIT=$(git -C BedrockItemUpgradeSchema rev-parse HEAD)
rm -rf BedrockBlockUpgradeSchema BedrockItemUpgradeSchema

git add blockupgrader/schemas itemupgrader/schemas
git commit -m "Updated upgrade schemas from https://github.com/pmmp/BedrockBlockUpgradeSchema/commit/$BLOCK_COMMIT and https://github.com/pmmp/BedrockItemUpgradeSchema/commit/$ITEM_COMMIT"