package customresource

const profileCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: profiles.cr.kanister.io
spec:
  group: cr.kanister.io
  names:
    kind: Profile
    listKind: ProfileList
    plural: profiles
    singular: profile
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          credential:
            properties:
              keyPair:
                properties:
                  idField:
                    type: string
                  secret:
                    properties:
                      apiVersion:
                        description: API version of the referent.
                        type: string
                      group:
                        description: API Group of the referent.
                        type: string
                      kind:
                        description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                        type: string
                      name:
                        description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: http://kubernetes.io/docs/user-guide/namespaces'
                        type: string
                      resource:
                        description: Resource name of the referent.
                        type: string
                    type: object
                  secretField:
                    type: string
                type: object
              secret:
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  group:
                    description: API Group of the referent.
                    type: string
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: http://kubernetes.io/docs/user-guide/namespaces'
                    type: string
                  resource:
                    description: Resource name of the referent.
                    type: string
                type: object
              type:
                type: string
            type: object
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          location:
            properties:
              bucket:
                type: string
              endpoint:
                type: string
              prefix:
                type: string
              region:
                type: string
              type:
                type: string
            type: object
          metadata:
            type: object
          skipSSLVerify:
            type: boolean
        type: object
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
