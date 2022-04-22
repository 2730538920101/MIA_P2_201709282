package structs

type MasterBootRecord struct{
	Mbr_tamano [100]byte 
	Mbr_fecha_creacion [100]byte
	Mbr_disk_signature [100]byte
	Disk_fit [100]byte
	Mbr_partition_1 Partition 
	Mbr_partition_2 Partition
	Mbr_partition_3 Partition
	Mbr_partition_4 Partition
}

func NewMBR(tam, date, sign, ft string) MasterBootRecord {
	e := MasterBootRecord{}
	copy(e.Mbr_tamano[:], tam)
	copy(e.Mbr_fecha_creacion[:], date)
	copy(e.Mbr_disk_signature[:], sign)
	copy(e.Disk_fit[:], ft)
	return e
}

type Partition struct{
	Part_status []byte
	Part_type []byte
	Part_fit []byte
	Part_start []byte
	Part_size []byte
	Part_name []byte
}

func NewPartition(stat, tipo, ft, st, tam, nom string) Partition{
	e := Partition{}
	copy(e.Part_status[:], stat)
	copy(e.Part_type[:], tipo)
	copy(e.Part_fit[:], ft)
	copy(e.Part_start[:], st)
	copy(e.Part_size[:], tam)
	copy(e.Part_name[:], nom)
	return e
}

type ExtendedBootRecord struct{
	Part_status []byte
	Part_fit []byte
	Part_start []byte
	Part_size []byte
	Part_next []byte
	Part_name []byte
}

type SuperBlock struct{
	S_filesystem_type []byte
	S_inodes_count []byte
	S_blocks_count []byte
	S_free_blocks_count []byte
	S_free_inodes_count []byte
	S_mtime []byte
	S_mnt_count []byte
	S_magic []byte
	S_inode_size []byte
	S_block_size []byte
	S_first_ino []byte
	S_first_blo []byte
	S_bm_inode_start []byte
	S_bm_block_start []byte
	S_inode_start []byte
	S_block_start []byte
}

type TablaInodos struct{
	I_uid []byte
	I_gid []byte
	I_size []byte
	I_atime []byte
	I_ctime []byte
	I_mtime []byte
	I_block []byte
	I_type []byte
	I_perm []byte
}

type BlockDirectory struct{
	B_content [4]Content
}

type Content struct{
	B_name []byte
	B_inodo []byte
}

type BlockFile struct{
	B_content []byte
}

