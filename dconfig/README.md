Prep
1. kubectl config use-context docker-desktop
2. kubectl get nodes (should see docker-desktop)


3. docker build -t demo:local .
2. kubectl apply -f configmap.yaml
3. kubectl apply -f deployment.yaml
4. kubectl logs -f deploy/demo


In a separate console: 
1. kubectl edit cm demo-config
2. watch for log update


Example Output:
2026/02/03 03:46:43 EVENT: CREATE     /etc/config/..2026_02_03_03_46_43.2002777651
2026/02/03 03:46:43 🔥 reload: version: v5

2026/02/03 03:46:43 EVENT: CHMOD      /etc/config/..2026_02_03_03_46_43.2002777651
2026/02/03 03:46:43 EVENT: CREATE     /etc/config/..data_tmp
2026/02/03 03:46:43 🔥 reload: version: v5

2026/02/03 03:46:43 EVENT: RENAME     /etc/config/..data_tmp
2026/02/03 03:46:44 🔥 reload: version: v5

2026/02/03 03:46:44 EVENT: CREATE     /etc/config/..data
2026/02/03 03:46:44 🔥 reload: version: v5

2026/02/03 03:46:44 EVENT: REMOVE     /etc/config/..2026_02_03_03_45_28.2102084789
2026/02/03 03:46:44 🔥 reload: version: v5