#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
API_ROOT="${SCRIPT_DIR}"
CORE_ROOT="${EMPTYDEA_CORE_ROOT:-$(cd "${SCRIPT_DIR}/../EmptyDea-core" && pwd)}"

PROTOCOL_SRC="${1:-${MOUSE_PROTOCOL_SRC:-/root/Yeah114/mousetunnel/minecraft/protocol}}"
PROTO_ROOT="${API_ROOT}/proto"
PROTOCOL_PROTO_OUT="${PROTO_ROOT}/minecraft/protocol"
PB_OUT="${API_ROOT}/pb"
CONVERTER_OUT="${CORE_ROOT}/frame/EmptyDeaCore/converter"

GO_PACKAGE_ROOT="github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
CONVERTER_PACKAGE_ROOT="github.com/Yeah114/EmptyDea-core/frame/EmptyDeaCore/converter"

if [[ ! -d "${PROTOCOL_SRC}" ]]; then
	echo "protocol source directory does not exist: ${PROTOCOL_SRC}" >&2
	exit 1
fi

export PATH="${PATH}:$(go env GOPATH)/bin"

rm -rf "${PROTOCOL_PROTO_OUT}" "${PB_OUT}" "${CONVERTER_OUT}"
mkdir -p "${PROTOCOL_PROTO_OUT}" "${PB_OUT}" "${CONVERTER_OUT}"

cd "${CORE_ROOT}"

go run ./cmd/protocol_proto_gen \
	-src "${PROTOCOL_SRC}" \
	-out "${PROTOCOL_PROTO_OUT}" \
	-server_out "${CONVERTER_OUT}" \
	-go_package_root "${GO_PACKAGE_ROOT}" \
	-server_package_root "${CONVERTER_PACKAGE_ROOT}"

mapfile -t PROTO_FILES < <(find "${PROTOCOL_PROTO_OUT}" -name '*.proto' | sort)
if [[ "${#PROTO_FILES[@]}" -eq 0 ]]; then
	echo "no proto files generated under ${PROTOCOL_PROTO_OUT}" >&2
	exit 1
fi

protoc \
	-I "${PROTOCOL_PROTO_OUT}" \
	--go_out="${API_ROOT}" \
	--go_opt=module=github.com/EmptyDea-Team/EmptyDea-core-api \
	"${PROTO_FILES[@]}"

mapfile -t API_PROTO_FILES < <(find "${PROTO_ROOT}" -name '*.proto' \
	! -path "${PROTOCOL_PROTO_OUT}/*" \
	| sort)
if [[ "${#API_PROTO_FILES[@]}" -gt 0 ]]; then
	protoc \
		-I "${PROTO_ROOT}" \
		-I "${PROTOCOL_PROTO_OUT}" \
		--go_out="${API_ROOT}" \
		--go_opt=module=github.com/EmptyDea-Team/EmptyDea-core-api \
		--go-grpc_out="${API_ROOT}" \
		--go-grpc_opt=module=github.com/EmptyDea-Team/EmptyDea-core-api \
		"${API_PROTO_FILES[@]}"
fi
