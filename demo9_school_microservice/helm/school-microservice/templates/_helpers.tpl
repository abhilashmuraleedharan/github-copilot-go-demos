# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-21
{{/*
Expand the name of the chart.
*/}}
{{- define "school-microservice.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
*/}}
{{- define "school-microservice.fullname" -}}
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
{{- define "school-microservice.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "school-microservice.labels" -}}
helm.sh/chart: {{ include "school-microservice.chart" . }}
{{ include "school-microservice.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "school-microservice.selectorLabels" -}}
app.kubernetes.io/name: {{ include "school-microservice.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "school-microservice.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "school-microservice.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Generate image name for a service
*/}}
{{- define "school-microservice.image" -}}
{{- $registry := .Values.global.imageRegistry | default .Values.image.registry -}}
{{- $repository := .service.image.repository -}}
{{- $tag := .Values.image.tag | default .Chart.AppVersion -}}
{{- if $registry }}
{{- printf "%s/%s/%s:%s" $registry .Values.image.repository $repository $tag }}
{{- else }}
{{- printf "%s/%s:%s" .Values.image.repository $repository $tag }}
{{- end }}
{{- end }}

{{/*
Generate service name
*/}}
{{- define "school-microservice.serviceName" -}}
{{- printf "%s-%s" (include "school-microservice.fullname" .) .serviceName }}
{{- end }}

{{/*
Generate common environment variables
*/}}
{{- define "school-microservice.commonEnvVars" -}}
- name: COUCHBASE_HOST
  valueFrom:
    configMapKeyRef:
      name: {{ include "school-microservice.fullname" . }}-config
      key: COUCHBASE_HOST
- name: COUCHBASE_BUCKET
  valueFrom:
    configMapKeyRef:
      name: {{ include "school-microservice.fullname" . }}-config
      key: COUCHBASE_BUCKET
- name: COUCHBASE_USERNAME
  valueFrom:
    secretKeyRef:
      name: {{ include "school-microservice.fullname" . }}-secret
      key: COUCHBASE_USERNAME
- name: COUCHBASE_PASSWORD
  valueFrom:
    secretKeyRef:
      name: {{ include "school-microservice.fullname" . }}-secret
      key: COUCHBASE_PASSWORD
{{- end }}
