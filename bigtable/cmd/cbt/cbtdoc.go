// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// DO NOT EDIT. THIS IS AUTOMATICALLY GENERATED.
// Run "go generate" to regenerate.
//go:generate go run cbt.go gcpolicy.go -o cbtdoc.go doc

/*
Cbt is a command-line tool that allows you to interact with Cloud Bigtable.
The `cbt` tool is a component of the Cloud SDK. You need to [install the Cloud SDK](https://cloud.google.com/bigtable/docs/cbt-overview)
and the `cbt` tool to use the following commands:

Usage:

    cbt [<var>option</var>] command <var>required-argument</var> [<var>optional-argument</var>]

The commands are:

    count                     Count rows in a table
    createinstance            Create an instance with an initial cluster
    createcluster             Create a cluster in the configured instance
    createfamily              Create a column family
    createtable               Create a table
    updatecluster             Update a cluster in the configured instance
    deleteinstance            Delete an instance
    deletecluster             Delete a cluster from the configured instance
    deletecolumn              Delete all cells in a column
    deletefamily              Delete a column family
    deleterow                 Delete a row
    deletetable               Delete a table
    doc                       Print godoc-suitable documentation for cbt
    help                      Print help text
    listinstances             List instances in a project
    listclusters              List clusters in an instance
    lookup                    Read from a single row
    ls                        List tables and column families
    mddoc                     Print documentation for cbt in Markdown format
    read                      Read rows
    set                       Set value of a cell (write)
    setgcpolicy               Set the garbage-collection policy (age, versions) for a column family
    waitforreplication        Block until all the completed writes have been replicated to all the clusters
    createtablefromsnapshot   Create a table from a snapshot (snapshots alpha)
    createsnapshot            Create a snapshot from a source table (snapshots alpha)
    listsnapshots             List snapshots in a cluster (snapshots alpha)
    getsnapshot               Get snapshot info (snapshots alpha)
    deletesnapshot            Delete snapshot in a cluster (snapshots alpha)
    version                   Print the current cbt version
    createappprofile          Creates app profile for an instance
    getappprofile             Reads app profile for an instance
    listappprofile            Lists app profile for an instance
    updateappprofile          Updates app profile for an instance
    deleteappprofile          Deletes app profile for an instance

The options are:

    -project string
        project ID. If unset uses gcloud configured project
    -instance string
        Cloud Bigtable instance
    -creds string
        Path to the credentials file. If set, uses the application credentials in this file

Use "cbt help \<command>" for more information about a command.


Alpha features are not currently available to most Cloud Bigtable customers. The
features might be changed in backward-incompatible ways and are not recommended
for production use. They are not subject to any SLA or deprecation policy.

Note: cbt does not support specifying arbitrary bytes on the command line for
any value that Cloud Bigtable otherwise supports (for example, the row key and
column qualifier).

For convenience, values of the -project, -instance, -creds,
-admin-endpoint and -data-endpoint flags may be specified in
~/.cbtrc in this format:

    project = my-project-123
    instance = my-instance
    creds = path-to-account-key.json
    admin-endpoint = hostname:port
    data-endpoint = hostname:port

All values are optional, and all will be overridden by flags.




Count rows in a table

Usage:
	cbt count <table>




Create an instance with an initial cluster

Usage:
	cbt createinstance <instance-id> <display-name> <cluster-id> <zone> <num-nodes> <storage type>
	  instance-id      Permanent, unique id for the instance
	  display-name     Description of the instance
	  cluster-id       Permanent, unique id for the cluster in the instance
	  zone             The zone in which to create the cluster
	  num-nodes        The number of nodes to create
	  storage-type     SSD or HDD

	    Example: cbt createinstance my-instance "My instance" my-instance-c1 us-central1-b 3 SSD




Create a cluster in the configured instance

Usage:
	cbt createcluster <cluster-id> <zone> <num-nodes> <storage type>
	  cluster-id       Permanent, unique id for the cluster in the instance
	  zone             The zone in which to create the cluster
	  num-nodes        The number of nodes to create
	  storage-type     SSD or HDD

	    Example: cbt createcluster my-instance-c2 europe-west1-b 3 SSD




Create a column family

Usage:
	cbt createfamily <table> <family>

	    Example cbt createfamily mobile-time-series stats_summary




Create a table

Usage:
	cbt createtable <table> [families=family[:gcpolicy],...] [splits=split,...]
	  families     Column families and their associated GC policies. Make sure to quote
	               this if using shell operators && and ||. For gcpolicy,
	               see "setgcpolicy".
	  splits       Row key to be used to initially split the table

	    Example: cbt createtable mobile-time-series "families=stats_summary:maxage=10d||maxversions=1,stats_detail:maxage=10d||maxversions=1" splits=tablet,phone




Update a cluster in the configured instance

Usage:
	cbt updatecluster <cluster-id> [num-nodes=num-nodes]
	  cluster-id    Permanent, unique id for the cluster in the instance
	  num-nodes     The new number of nodes

	    Example: cbt updatecluster my-instance-c1 num-nodes=5




Delete an instance

Usage:
	cbt deleteinstance <instance>

	    Example: cbt deleteinstance my-instance




Delete a cluster from the configured instance

Usage:
	cbt deletecluster <cluster>

	    Example: cbt deletecluster my-instance-c2




Delete all cells in a column

Usage:
	cbt deletecolumn <table> <row> <family> <column> [app-profile=<app profile id>]
	  app-profile=<app profile id>        The app profile id to use for the request

	    Example: cbt deletecolumn mobile-time-series phone#4c410523#20190501 stats_summary os_name




Delete a column family

Usage:
	cbt deletefamily <table> <family>

	    Example: cbt deletefamily mobile-time-series stats_summary




Delete a row

Usage:
	cbt deleterow <table> <row> [app-profile=<app profile id>]
	  app-profile=<app profile id>        The app profile id to use for the request

	    Example: cbt deleterow mobile-time-series phone#4c410523#20190501




Delete a table

Usage:
	cbt deletetable <table>

	    Example: cbt deletetable mobile-time-series




Print godoc-suitable documentation for cbt

Usage:
	cbt doc




Print help text

Usage:
	cbt help [command]

	    Example: cbt help createtable




List instances in a project

Usage:
	cbt listinstances




List clusters in an instance

Usage:
	cbt listclusters




Read from a single row

Usage:
	cbt lookup <table> <row> [columns=[family]:[qualifier],...] [cells-per-column=<n>]  [app-profile=<app profile id>]
	  columns=[family]:[qualifier],...    Read only these columns, comma-separated
	  cells-per-column=<n>                Read only this many cells per column
	  app-profile=<app profile id>        The app profile id to use for the request

	 Example: cbt lookup mobile-time-series phone#4c410523#20190501 columns=stats_summary:os_build,os_name cells-per-column=1




List tables and column families

Usage:
	cbt ls                List tables
	cbt ls <table>        List column families in <table>

	    Example: cbt ls mobile-time-series




Print documentation for cbt in Markdown format

Usage:
	cbt mddoc




Read rows

Usage:
	cbt read <table> [start=<row>] [end=<row>] [prefix=<prefix>] [regex=<regex>] [columns=[family]:[qualifier],...] [count=<n>] [cells-per-column=<n>] [app-profile=<app profile id>]
	  start=<row>                         Start reading at this row
	  end=<row>                           Stop reading before this row
	  prefix=<prefix>                     Read rows with this prefix
	  regex=<regex>                       Read rows with keys matching this regex
	  columns=[family]:[qualifier],...    Read only these columns, comma-separated
	  count=<n>                           Read only this many rows
	  cells-per-column=<n>                Read only this many cells per column
	  app-profile=<app profile id>        The app profile id to use for the request

	    Examples (see write row, to have data to read):
	      cbt read mobile-time-series prefix=phone columns=stats_summary:os_build,os_name count=10
	      cbt read mobile-time-series start=phone#4c410523#20190501 end=phone#4c410523#20190601
	      cbt read mobile-time-series regex="phone.*" cells-per-column=1





Set value of a cell (write)

Usage:
	cbt set <table> <row> [app-profile=<app profile id>] family:column=val[@ts] ...
	  app-profile=<app profile id>        The app profile id to use for the request
	  family:column=val[@ts] may be repeated to set multiple cells.

	    ts is an optional integer timestamp.
	    If it cannot be parsed, the `@ts` part will be
	    interpreted as part of the value.

	    Examples:
	      cbt set mobile-time-series phone#4c410523#20190501 stats_summary:connected_cell=1@12345 stats_summary:connected_cell=0@54321
	      cbt set mobile-time-series phone#4c410523#20190501 stats_summary:os_build=PQ2A.190405.003 stats_summary:os_name=android




Set the garbage-collection policy (age, versions) for a column family

Usage:
	cbt setgcpolicy <table> <family> ((maxage=<d> | maxversions=<n>) [(and|or) (maxage=<d> | maxversions=<n>),...] | never)
	  maxage=<d>         Maximum timestamp age to preserve (e.g. "1h", "4d")
	  maxversions=<n>    Maximum number of versions to preserve

	    Examples:
	      cbt setgcpolicy mobile-time-series stats_detail maxage=10d
	      cbt setgcpolicy mobile-time-series stats_summary maxage=10d or maxversion=1





Block until all the completed writes have been replicated to all the clusters

Usage:
	cbt waitforreplication <table>




Create a table from a snapshot (snapshots alpha)

Usage:
	cbt createtablefromsnapshot <table> <cluster> <snapshot>
	  table        The name of the table to create
	  cluster      The cluster where the snapshot is located
	  snapshot     The snapshot to restore





Create a snapshot from a source table (snapshots alpha)

Usage:
	cbt createsnapshot <cluster> <snapshot> <table> [ttl=<d>]
	  [ttl=<d>]        Lifespan of the snapshot (e.g. "1h", "4d")




List snapshots in a cluster (snapshots alpha)

Usage:
	cbt listsnapshots [<cluster>]




Get snapshot info (snapshots alpha)

Usage:
	cbt getsnapshot <cluster> <snapshot>




Delete snapshot in a cluster (snapshots alpha)

Usage:
	cbt deletesnapshot <cluster> <snapshot>




Print the current cbt version

Usage:
	cbt version




Creates app profile for an instance

Usage:
	usage: cbt createappprofile <instance-id> <profile-id> <description> (route-any | [ route-to=<cluster-id> : transactional-writes]) [optional flag]
	optional flags may be `force`

	    Examples:
	      cbt createappprofile my-instance multi-cluster "Routes to nearest available cluster" route-any
	      cbt createappprofile my-instance single-cluster "Europe routing" route-to=my-instance-c2




Reads app profile for an instance

Usage:
	cbt getappprofile <instance-id> <profile-id>




Lists app profile for an instance

Usage:
	cbt listappprofile <instance-id>




Updates app profile for an instance

Usage:
	usage: cbt updateappprofile  <instance-id> <profile-id> <description>(route-any | [ route-to=<cluster-id> : transactional-writes]) [optional flag]
	optional flags may be `force`

	    Example: cbt updateappprofile my-instance multi-cluster "Use this one." route-any




Deletes app profile for an instance

Usage:
	cbt deleteappprofile <instance-id> <profile-id>

	    Example: cbt deleteappprofile my-instance single-cluster




*/
package main
