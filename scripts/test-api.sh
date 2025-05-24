#!/bin/sh

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Use environment variable or default
API_URL=${API_URL:-"http://api:8080"}
TEST_EMAIL="test$(date +%s)@example.com"
TEST_PASSWORD="password123"
TEST_NAME="Test User"

echo -e "${YELLOW}🚀 Starting API Integration Tests${NC}"
echo "API URL: $API_URL"
echo "=================================="

# Test 1: Wait for API to be ready
echo -e "\n${YELLOW}Test 1: Waiting for API to be ready${NC}"
sleep 5

# Test 2: User Registration
echo -e "\n${YELLOW}Test 2: User Registration${NC}"
REGISTER_RESPONSE=$(curl -s -X POST $API_URL/auth/register \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"$TEST_NAME\",\"email\":\"$TEST_EMAIL\",\"password\":\"$TEST_PASSWORD\"}")

echo "Register Response: $REGISTER_RESPONSE"

if echo "$REGISTER_RESPONSE" | grep -q "successfully\|User\|created"; then
    echo -e "${GREEN}✅ User registration successful${NC}"
else
    echo -e "${RED}❌ User registration failed${NC}"
    exit 1
fi

# Test 3: User Login
echo -e "\n${YELLOW}Test 3: User Login${NC}"
LOGIN_RESPONSE=$(curl -s -X POST $API_URL/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$TEST_EMAIL\",\"password\":\"$TEST_PASSWORD\"}")

echo "Login Response: $LOGIN_RESPONSE"

# Extract token
TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ ! -z "$TOKEN" ] && [ "$TOKEN" != "null" ]; then
    echo -e "${GREEN}✅ User login successful${NC}"
    echo "Token: ${TOKEN:0:20}..."
else
    echo -e "${RED}❌ User login failed${NC}"
    exit 1
fi

# Test 4: Protected Endpoint
echo -e "\n${YELLOW}Test 4: Protected Endpoint Access${NC}"
USERS_RESPONSE=$(curl -s -X GET $API_URL/users \
  -H "Authorization: Bearer $TOKEN")

echo "Users Response: $USERS_RESPONSE"

if echo "$USERS_RESPONSE" | grep -q "users\|\[\]"; then
    echo -e "${GREEN}✅ Protected endpoint accessible${NC}"
else
    echo -e "${RED}❌ Protected endpoint failed${NC}"
    exit 1
fi

echo -e "\n${GREEN}🎉 All tests completed successfully!${NC}"