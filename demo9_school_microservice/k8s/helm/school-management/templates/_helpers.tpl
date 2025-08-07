{{/*
Expand the name of the chart.
*/}}
{{- define "school-management.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "school-management.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "school-management.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "school-management.labels" -}}
helm.sh/chart: {{ include "school-management.chart" . }}
{{ include "school-management.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "school-management.selectorLabels" -}}
app.kubernetes.io/name: {{ include "school-management.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "school-management.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "school-management.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
API Gateway labels
*/}}
{{- define "school-management.apiGateway.labels" -}}
{{ include "school-management.labels" . }}
app.kubernetes.io/component: api-gateway
{{- end }}

{{/*
API Gateway selector labels
*/}}
{{- define "school-management.apiGateway.selectorLabels" -}}
{{ include "school-management.selectorLabels" . }}
app.kubernetes.io/component: api-gateway
{{- end }}

{{/*
Student Service labels
*/}}
{{- define "school-management.studentService.labels" -}}
{{ include "school-management.labels" . }}
app.kubernetes.io/component: student-service
{{- end }}

{{/*
Student Service selector labels
*/}}
{{- define "school-management.studentService.selectorLabels" -}}
{{ include "school-management.selectorLabels" . }}
app.kubernetes.io/component: student-service
{{- end }}

{{/*
Teacher Service labels
*/}}
{{- define "school-management.teacherService.labels" -}}
{{ include "school-management.labels" . }}
app.kubernetes.io/component: teacher-service
{{- end }}

{{/*
Teacher Service selector labels
*/}}
{{- define "school-management.teacherService.selectorLabels" -}}
{{ include "school-management.selectorLabels" . }}
app.kubernetes.io/component: teacher-service
{{- end }}

{{/*
Academic Service labels
*/}}
{{- define "school-management.academicService.labels" -}}
{{ include "school-management.labels" . }}
app.kubernetes.io/component: academic-service
{{- end }}

{{/*
Academic Service selector labels
*/}}
{{- define "school-management.academicService.selectorLabels" -}}
{{ include "school-management.selectorLabels" . }}
app.kubernetes.io/component: academic-service
{{- end }}

{{/*
Achievement Service labels
*/}}
{{- define "school-management.achievementService.labels" -}}
{{ include "school-management.labels" . }}
app.kubernetes.io/component: achievement-service
{{- end }}

{{/*
Achievement Service selector labels
*/}}
{{- define "school-management.achievementService.selectorLabels" -}}
{{ include "school-management.selectorLabels" . }}
app.kubernetes.io/component: achievement-service
{{- end }}

{{/*
Couchbase labels
*/}}
{{- define "school-management.couchbase.labels" -}}
{{ include "school-management.labels" . }}
app.kubernetes.io/component: couchbase
{{- end }}

{{/*
Couchbase selector labels
*/}}
{{- define "school-management.couchbase.selectorLabels" -}}
{{ include "school-management.selectorLabels" . }}
app.kubernetes.io/component: couchbase
{{- end }}

{{/*
Generate image name
*/}}
{{- define "school-management.image" -}}
{{- $registry := .Values.global.imageRegistry -}}
{{- $repository := .repository -}}
{{- $tag := .Values.global.imageTag -}}
{{- printf "%s/%s:%s" $registry $repository $tag -}}
{{- end }}

{{/*
Database connection string
*/}}
{{- define "school-management.databaseUrl" -}}
{{- if .Values.couchbase.enabled -}}
{{- printf "couchbase://%s:%d" (include "school-management.fullname" .) .Values.couchbase.service.ports.data -}}
{{- else -}}
{{- printf "couchbase://localhost:11210" -}}
{{- end -}}
{{- end }}
