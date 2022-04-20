package structs

type MasterBootRecord struct{
	mbr_tamano []byte 
	mbr_fecha_creacion []byte
	mbr_disk_signature []byte
	disk_fit []byte
	mbr_partition_1 Partition 
	mbr_partition_2 Partition
	mbr_partition_3 Partition
	mbr_partition_4 Partition
}

type Partition struct{
	part_status []byte
	part_type []byte
	part_fit []byte
	part_start []byte
	part_size []byte
	part_name []byte
}

type ExtendedBootRecord struct{
	part_status []byte
	part_fit []byte
	part_start []byte
	part_size []byte
	part_next []byte
	part_name []byte
}

type SuperBlock struct{
	s_filesystem_type []byte
	s_inodes_count []byte
	s_blocks_count []byte
	s_free_blocks_count []byte
	s_free_inodes_count []byte
	s_mtime []byte
	s_mnt_count []byte
	s_magic []byte
	s_inode_size []byte
	s_block_size []byte
	s_first_ino []byte
	s_first_blo []byte
	s_bm_inode_start []byte
	s_bm_block_start []byte
	s_inode_start []byte
	s_block_start []byte
}

type TablaInodos struct{
	i_uid []byte
	i_gid []byte
	i_size []byte
	i_atime []byte
	i_ctime []byte
	i_mtime []byte
	i_block []byte
	i_type []byte
	i_perm []byte
}

type BlockDirectory struct{
	b_content [4]Content
}

type Content struct{
	b_name []byte
	b_inodo []byte
}

type BlockFile struct{
	b_content []byte
}