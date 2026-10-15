#!/usr/bin/env bash

set -euo pipefail

image=postgres:18.6
host=${PGHOST:-host.docker.internal}
port=${PGPORT:-5432}
user=${PGUSER:-postgres}
database=${PGDATABASE:-postgres}
password=${PGPASSWORD:-}

usage() {
	printf 'Usage: %s [-h host] [-p port] [-U user] [-d database] [-W password] [sql-file]\n' "${0##*/}" >&2
}

while getopts ':h:p:U:d:W:' option; do
	case "$option" in
		h) host=$OPTARG ;;
		p) port=$OPTARG ;;
		U) user=$OPTARG ;;
		d) database=$OPTARG ;;
		W) password=$OPTARG ;;
		:|\?) usage; exit 2 ;;
	esac
done
shift "$((OPTIND - 1))"

if (($# > 1)); then
	usage
	exit 2
fi

case "$host" in
	localhost|127.0.0.1|::1) host=host.docker.internal ;;
esac

docker_args=(run --rm -i)
if (($# == 0)) && [[ -t 0 ]]; then
	docker_args+=(-t)
fi
docker_args+=(--add-host=host.docker.internal:host-gateway)
if [[ -n "$password" ]]; then
	docker_args+=(-e "PGPASSWORD=$password")
fi

if (($# == 1)); then
	sql_file=$1
	if [[ ! -f "$sql_file" ]]; then
		printf 'SQL file %q is not a regular file\n' "$sql_file" >&2
		exit 1
	fi
	sql_file=$(cd -- "$(dirname -- "$sql_file")" && printf '%s/%s\n' "$PWD" "$(basename -- "$sql_file")")
	docker_args+=(-v "$sql_file:/query.sql:ro")
fi

docker_args+=("$image" psql)
[[ -n "$host" ]] && docker_args+=(-h "$host")
[[ -n "$port" ]] && docker_args+=(-p "$port")
[[ -n "$user" ]] && docker_args+=(-U "$user")
[[ -n "$database" ]] && docker_args+=(-d "$database")
if (($# == 1)); then
	docker_args+=(-f /query.sql)
fi

exec docker "${docker_args[@]}"
