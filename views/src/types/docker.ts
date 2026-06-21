export interface DockerContainer {
	Id: string;
	Names: string[];
	Image: string;
	ImageID: string;
	ImageManifestDescriptor: ImageManifestDescriptor;
	Command: string;
	Created: number;
	Ports: Port[];
	Labels: Record<string, string>;
	State: string;
	Status: string;
	HostConfig: HostConfig;
	Health: Health;
	NetworkSettings: NetworkSettings;
	Mounts: Mount[];
}

export interface ImageManifestDescriptor {
	mediaType: string;
	digest: string;
	size: number;
	annotations: Record<string, string>;
	platform: Platform;
}

export interface Platform {
	architecture: string;
	os: string;
}

export interface Port {
	IP?: string;
	PrivatePort?: number;
	PublicPort?: number;
	Type?: string;
}

export interface HostConfig {
	NetworkMode: string;
}

export interface Health {
	Status: string;
	FailingStreak: number;
}

export interface NetworkSettings {
	Networks: Record<string, Network>;
}

export interface Network {
	IPAMConfig: unknown | null;
	Links: unknown | null;
	Aliases: string[] | null;
	DriverOpts: Record<string, string> | null;
	GwPriority: number;
	NetworkID: string;
	EndpointID: string;
	Gateway: string;
	IPAddress: string;
	MacAddress: string;
	IPPrefixLen: number;
	IPv6Gateway: string;
	GlobalIPv6Address: string;
	GlobalIPv6PrefixLen: number;
	DNSNames: string[] | null;
}

export interface Mount {
	Type?: string;
	Source?: string;
	Destination?: string;
	Mode?: string;
	RW?: boolean;
	Propagation?: string;
}
