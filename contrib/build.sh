#!/bin/sh

# run me to do a multi-arch build and deploy packages directly to ghcr.io!

ARCHS=(arm64 arm amd64)
IMAGE_NAME=ghcr.io/hugomd/cloudflare-ddns

OCI_CONTAINER_BUILDER=podman
#OCI_CONTAINER_BUILDER=docker

$OCI_CONTAINER_BUILDER image rm $IMAGE_NAME
$OCI_CONTAINER_BUILDER manifest create $IMAGE_NAME

for arch in ${ARCHS[@]}; do
    $OCI_CONTAINER_BUILDER build --build-arg "ARCH=$arch" .. -f Dockerfile -t $IMAGE_NAME:$arch
    $OCI_CONTAINER_BUILDER push $IMAGE_NAME:$arch
    $OCI_CONTAINER_BUILDER manifest add --arch $arch --os linux $IMAGE_NAME $IMAGE_NAME:$arch
done

$OCI_CONTAINER_BUILDER manifest push --format v2s2 --all $IMAGE_NAME docker://$IMAGE_NAME:latest
$OCI_CONTAINER_BUILDER manifest push --format v2s2 --all --purge $IMAGE_NAME docker://$IMAGE_NAME:$VERSION
echo "done"
