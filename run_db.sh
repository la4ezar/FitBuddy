POSTGRES_CONTAINER="test-postgres"
POSTGRES_VERSION="15"

DB_USER="postgres"
DB_PWD="pgsql@12345"
DB_NAME="fitbuddy"
DB_PORT="5432"
DB_HOST="127.0.0.1"

echo -e "${GREEN}Start Postgres in detached mode${NC}"
docker run -d --name ${POSTGRES_CONTAINER} \
            -e POSTGRES_HOST=${DB_HOST} \
            -e POSTGRES_USER=${DB_USER} \
            -e POSTGRES_PASSWORD=${DB_PWD} \
            -e POSTGRES_DB=${DB_NAME} \
            -e POSTGRES_PORT=${DB_PORT} \
            -p ${DB_PORT}:${DB_PORT} \
            postgres:${POSTGRES_VERSION}

echo '# WAITING FOR CONNECTION WITH DATABASE #'
for i in {1..30}
do
    docker exec ${POSTGRES_CONTAINER} pg_isready -U "${DB_USER}" -h "${DB_HOST}" -p "${DB_PORT}" -d "${DB_NAME}"
    if [ $? -eq 0 ]
    then
        dbReady=true
        break
    fi
    sleep 1
done

if [ "${dbReady}" != true ] ; then
    echo '# COULD NOT ESTABLISH CONNECTION TO DATABASE #'
    exit 1
fi
