# 简介
* 基于kubeflow pipeline 0.1.31版本做的二次开发。
* 修改界面
* 创建数据库表
* 开发新的api
* 开发新的pod、service
* 我的blog有相关文章：https://blog.csdn.net/reachyu

# 修改界面文件列表
* frontend\src\components\SideNav.tsx
* frontend\src\components\Banner.tsx
* frontend\src\components\UploadPipelineDialog.tsx
* frontend\src\components\CustomTable.tsx
* frontend\src\components\NewRunParameters.tsx
* frontend\src\pages\ExperimentList.tsx
* frontend\src\pages\ExperimentsAndRuns.tsx
* frontend\src\pages\Page.tsx
* frontend\src\pages\NewRun.tsx
* frontend\src\pages\Archive.tsx
* frontend\src\pages\PipelineList.tsx
* frontend\src\pages\RunList.tsx
* frontend\src\pages\RunDetails.tsx
* frontend\src\pages\NewExperiment.tsx
* frontend\public\index.html
* frontend\public\manifest.json
* frontend\src\lib\Buttons.ts

# 创建数据库表文件列表
* backend\src\apiserver\model\BUILD.bazel
* backend\src\apiserver\model\deepexi_app.go
* backend\src\apiserver\model\deepexi_dataset.go
* backend\src\apiserver\client_manager.go

# 开发新的api文件列表
* backend\src\apiserver\client_manager.go
* backend\src\apiserver\main.go
* backend\src\apiserver\storage\myapp_manage_store.go
* backend\src\apiserver\resource\resource_manager.go
* backend\src\apiserver\resource\client_manager_fake.go
* backend\src\apiserver\server\myapp_manage_server.go
* backend\src\apiserver\server\BUILD.bazel
* backend\src\apiserver\storage\BUILD.bazel


# 效果
* 修改前的页面  
![image](https://img-blog.csdnimg.cn/20200625094923201.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3JlYWNoeXU=,size_16,color_FFFFFF,t_70)
* 修改后的页面  
![image](https://img-blog.csdnimg.cn/20200625094730671.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3JlYWNoeXU=,size_16,color_FFFFFF,t_70)

