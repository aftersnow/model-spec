# CNCF ModelPack Specification Standard

[![GoDoc](https://godoc.org/github.com/modelpack/model-spec?status.svg)](https://godoc.org/github.com/modelpack/model-spec)
[![Discussions](https://img.shields.io/badge/discussions-on%20github-blue?style=flat-square)](https://github.com/modelpack/model-spec/discussions)

The Cloud Native Computing Foundation's (CNCF) ModelPack project is a vendor-neutral, open source specification standard to package, distribute and run AI models in a cloud native environments. It's goal is to enable the creation of standard-compliant implementations that would move AI/ML project artifacts out of vendor-controlled, proprietary formats and into a standardized and interchangeable format that is compatible with the cloud-native ecosystem.

## Rationale

Looking back in history, there are clear trends in the evolution of infrastructure. At first, there is the machine centric infrastructure age. GNU/Linux was born there and we saw a boom of Linux distributions then. Then comes the Virtual Machine centric infrastructure age, where we saw the rise of cloud computing and the development of virtualization technologies. The third age is the container centric infrastructure, and we saw the rise of container technologies like Docker and Kubernetes. The fourth age, which has just begun, is the AI model centric infrastructure age, where we will see a burst of technologies and projects around AI model development and deployment.

![img](docs/img/infra-trends.png)

Each of the new ages has brought new technologies and new ways of thinking. The container centric infrastructure has brought us the OCI image specification, which has become the standard for packaging and distributing software. The AI model centric infrastructure will bring us new ways of packaging and distributing AI models. This model specification is an attempt to define a standard that aligns with the container standards that organizations and individuals have successfully relied on for the last decade.

## Current Work

This specification provides a compatible way to package and distribute models based on the current [OCI image specification](https://github.com/opencontainers/image-spec/) and [the artifacts guidelines](https://github.com/opencontainers/image-spec/blob/main/manifest.md#guidelines-for-artifact-usage). For compatibility reasons, it only contains part of the model metadata, and handles model artifacts as opaque binaries. However, it provides a convenient way to package AI models in the container image format and can be used as [OCI volume sources](https://github.com/kubernetes/enhancements/issues/4639) in Kubernetes environments.

For details, please see [the specification](docs/spec.md).

## Usage Examples

### Creating a Model Configuration

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
    
    v1 "github.com/modelpack/model-spec/specs-go/v1"
    "github.com/opencontainers/go-digest"
)

func main() {
    // Create a model configuration
    model := v1.Model{
        Descriptor: v1.ModelDescriptor{
            Name:        "llama3-8b-instruct",
            Family:      "llama3",
            Version:     "3.1",
            Authors:     []string{"meta-llama@meta.com"},
            Title:       "Llama 3 8B Instruct",
            Description: "Llama 3 is a large language model developed by Meta.",
            DocURL:      "https://llama.meta.com/",
            Licenses:    []string{"Apache-2.0"},
            CreatedAt:   &time.Time{},
        },
        Config: v1.ModelConfig{
            Architecture: v1.ArchitectureTransformer,
            Format:       v1.FormatPyTorch,
            ParamSize:    "8b",
            Precision:    v1.PrecisionFP16,
            Capabilities: &v1.ModelCapabilities{
                InputTypes:  []v1.Modality{v1.TextModality},
                OutputTypes: []v1.Modality{v1.TextModality},
                Reasoning:   &[]bool{true}[0],
                ToolUsage:   &[]bool{true}[0],
            },
        },
        ModelFS: v1.ModelFS{
            Type: "layers",
            DiffIDs: []digest.Digest{
                digest.FromString("model-weights-layer"),
            },
        },
    }
    
    // Marshal to JSON
    jsonData, _ := json.MarshalIndent(model, "", "  ")
    fmt.Println(string(jsonData))
}
```

### Validating a Model Configuration

```go
package main

import (
    "bytes"
    "log"
    
    "github.com/modelpack/model-spec/schema"
)

func main() {
    // Your model configuration JSON
    configJSON := []byte(`{
        "descriptor": {
            "name": "my-model",
            "version": "1.0"
        },
        "config": {
            "architecture": "transformer",
            "format": "onnx"
        },
        "modelfs": {
            "type": "layers",
            "diffIds": ["sha256:..."]
        }
    }`)
    
    // Validate the configuration
    validator := schema.ValidatorMediaTypeModelConfig
    err := validator.Validate(bytes.NewReader(configJSON))
    if err != nil {
        log.Fatalf("Validation failed: %v", err)
    }
    
    log.Println("Model configuration is valid!")
}
```

## LICENSE

Apache 2.0 License. Please see [LICENSE](LICENSE) for more information.

## Community, Support, Discussion

You can engage with this project by joining the discussion on our Slack channel: [#modelpack](https://cloud-native.slack.com/archives/C07T0V480LF) in the [CNCF Slack workspace](https://slack.cncf.io/).

This project holds inclusivity, empathy, and responsibility at our core. We follow the CNCF's [Code of Conduct](./code-of-conduct.md), you can read it to understand the values guiding our community.

The rules governing this project can be found in the [Governance policy document](./GOVERNANCE.md)

## Contributing

Any feedback, suggestions, and contributions are welcome. Please feel free to open an issue or pull request.

Especially, we look forward to integrating the model specification with different model registry implementations (like [Harbor](https://goharbor.io/) and [Kubeflow model registry](https://www.kubeflow.org/docs/components/model-registry/overview/)), as well as existing model centric infrastructure projects like [Huggingface](https://huggingface.co/), [KitOps](https://kitops.ml/), [Kubeflow](https://www.kubeflow.org/), [Lepton](https://www.lepton.ai/), [Ollama](https://github.com/ollama/ollama), [ORAS](https://oras.land/), and others.
