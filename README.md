# 📜 Updated README – Product Service on Azure Container Apps

# 🛠 Full Command Set to Deploy Go App to Azure Container Apps

## 1. 🔹 Set Environment Variables

```bash
ACR_NAME="azurecontainerdemo"
RESOURCE_GROUP="demo-azure-container-apps"
LOCATION="centralus"
APP_NAME="demoappchmaps2025"
ENV="demoappchmaps2025"
```

---

## 2. 🔹 Create Resource Group

```bash
az group create --name $RESOURCE_GROUP --location $LOCATION
```

---

## 3. 🔹 Create Azure Container Registry (ACR)

```bash
az acr create --name $ACR_NAME --resource-group $RESOURCE_GROUP --sku Basic --location $LOCATION
```

---

## 4. 🔹 Build Docker Image (Important for Mac M1/M2 Users)

### ⚠️ Special Step for Mac M1 / M2 (ARM64 processors)

Use **Docker Buildx** to build a Linux AMD64 compatible image:

```bash
# Create and use buildx builder
docker buildx create --use

# Build for linux/amd64 and push directly to ACR
docker buildx build --platform linux/amd64 -t $ACR_NAME.azurecr.io/product-service:latest --push .
```

✅ This ensures your container is compatible with Azure Container Apps (linux/amd64 architecture).

---

## 5. 🔹 (If not pushed already) Login to Azure Container Registry

```bash
az acr login --name $ACR_NAME
```

---

## 6. 🔹 Create Azure Container Apps Environment

```bash
az containerapp env create \
  --name $ENV \
  --resource-group $RESOURCE_GROUP \
  --location $LOCATION
```

---

## 7. 🔹 Deploy the Container App

```bash
az containerapp create \
  --name $APP_NAME \
  --resource-group $RESOURCE_GROUP \
  --environment $ENV \
  --image $ACR_NAME.azurecr.io/product-service:latest \
  --target-port 8080 \
  --ingress external \
  --registry-server $ACR_NAME.azurecr.io \
  --registry-username "<your-acr-username>" \
  --registry-password "<your-acr-password>"
```

✅ Get your ACR username and password via:

```bash
az acr credential show --name $ACR_NAME
```

Use the `username` and `password` values from the output.

---

## 8. 🔹 Get the Public URL of your Container App

```bash
az containerapp show \
  --name $APP_NAME \
  --resource-group $RESOURCE_GROUP \
  --query properties.configuration.ingress.fqdn \
  --output tsv
```

✅ Your public URL will look like:

```
https://demoappchmaps2025.centralus.azurecontainerapps.io
```

---

# ✅ Quick Access Endpoints:

- `https://your-app-url/health` → Health check
- `https://your-app-url/api/products` → Product APIs
- `https://your-app-url/swagger/index.html` → Swagger UI

---

# ⚡ Quick Notes:

- Always **build using `linux/amd64`** if you're using Mac M1/M2 to avoid OS/Architecture mismatch.
- Azure Container Apps **auto-scales** based on HTTP traffic.
- ACA will **scale to zero** if no traffic = saves cost automatically.

---

## Running Unit Tests

To run the unit tests for this project, use the following command:

```bash
go test ./...
```

This command recursively tests all packages in the current directory and its subdirectories. `go test` works by compiling and running test functions found in files ending with `_test.go`. Test functions must be named starting with `Test` and take a `*testing.T` argument.

## Checking Code Coverage

Code coverage is a metric that measures the percentage of your code that is exercised by your tests. To check the code coverage of your unit tests, you can use the following commands:

1.  **Generate a coverage profile:**

    ```bash
    go test ./... -coverprofile=coverage.out
    ```

    The `-coverprofile` flag creates a file named `coverage.out` that contains the coverage data.

2.  **View the coverage report in your browser:**

    ```bash
    go tool cover -html=coverage.out
    ```

    This command opens a new tab in your browser with a detailed HTML report. The report shows which lines of your code are covered by the tests (in green), and which are not (in red). This helps you identify which parts of your code are not being tested.

### Enforcing Code Coverage in CI/CD

In a CI/CD pipeline, you can enforce a minimum code coverage percentage to ensure that new code is adequately tested. The following command will fail the build if the total coverage is less than 80%:

```bash
go test ./... -cover -covermode=atomic -coverprofile=coverage.out && go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/\%//' | awk '{if ($1 < 80) {exit 1}}'
```

**Explanation of the command:**

*   `go test ./... -cover -covermode=atomic -coverprofile=coverage.out`: This runs the tests and generates a coverage profile. The `-covermode=atomic` flag is necessary for accurate coverage measurement in multi-threaded tests.
*   `go tool cover -func=coverage.out`: This prints the coverage percentage for each function.
*   `grep total`: This filters the output to show only the total coverage.
*   `awk '{print $3}'`: This extracts the coverage percentage.
*   `sed 's/\%//'`: This removes the '%' sign.
*   `awk '{if ($1 < 80) {exit 1}}'`: This checks if the coverage is less than 80%. If it is, the command exits with a non-zero status code, which will cause the CI/CD build to fail.

# 🏁 That's it!

Your **GoLang Product Service API** is now successfully running serverlessly on **Azure Container Apps**! 🚀

---