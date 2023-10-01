set -e

tag=$(date "+%y.%m.%d.%s")
git tag $tag
git push git $tag
