    if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		input.InferenceProfileIdentifier = aws.String(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}