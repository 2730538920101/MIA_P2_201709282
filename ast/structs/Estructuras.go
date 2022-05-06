package structs

import(
	"strconv"
)

type MasterBootRecord struct{
	Mbr_tamano [100]byte
	Mbr_fecha_creacion [100]byte
	Mbr_disk_signature [100]byte
	Disk_fit [100]byte
	Mbr_partition [4]Partition 
	
}

func NewMBR(tam, date, sign, ft string) MasterBootRecord {
	e := MasterBootRecord{}
	copy(e.Mbr_tamano[:], tam)
	copy(e.Mbr_fecha_creacion[:], date)
	copy(e.Mbr_disk_signature[:], sign)
	copy(e.Disk_fit[:], ft)
	for i := 0; i < 4; i++{
		copy(e.Mbr_partition[i].Part_status[:],"INACTIVO")
		copy(e.Mbr_partition[i].Part_type[:],"P")
		copy(e.Mbr_partition[i].Part_fit[:],ft)
		copy(e.Mbr_partition[i].Part_start[:],"0")
		copy(e.Mbr_partition[i].Part_size[:],"0")
		copy(e.Mbr_partition[i].Part_name[:],"PART"+strconv.Itoa(i+1))
	}
	return e
}

type Partition struct{
	Part_status [500]byte
	Part_type [500]byte
	Part_fit [500]byte
	Part_start [500]byte
	Part_size [500]byte
	Part_name [500]byte
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
	Part_status [100]byte
	Part_fit [100]byte
	Part_start [100]byte
	Part_size [100]byte
	Part_next [100]byte
	Part_name [100]byte
}

func NewEBR(stat, ft, st, tam, sig, nom string) ExtendedBootRecord{
	e := ExtendedBootRecord{}
	copy(e.Part_status[:], stat)
	copy(e.Part_fit[:], ft)
	copy(e.Part_start[:], st)
	copy(e.Part_size[:], tam)
	copy(e.Part_next[:], sig)
	copy(e.Part_name[:], nom)
	return e
}

type SuperBlock struct{
	S_filesystem_type [10000]byte
	S_inodes_count [10000]byte
	S_blocks_count [10000]byte
	S_free_blocks_count [10000]byte
	S_free_inodes_count [10000]byte
	S_mtime [10000]byte
	S_mnt_count [10000]byte
	S_magic [10000]byte
	S_inode_size [10000]byte
	S_block_size [10000]byte
	S_first_ino [10000]byte
	S_first_blo [10000]byte
	S_bm_inode_start [10000]byte
	S_bm_block_start [10000]byte
	S_inode_start [10000]byte
	S_block_start [10000]byte
}

func NewSuperBlock(tipofs, contInodos, fcontInodos, contBlocks, fcontBlocks, mtime, mnt_cont, magic, inodeTam, blockTam, primIno, primBlock, bm_inodest, bm_blockst, inodest, blockst string) SuperBlock{
	e := SuperBlock{}
	copy(e.S_filesystem_type[:], tipofs)
	copy(e.S_inodes_count[:], contInodos)
	copy(e.S_blocks_count[:], contBlocks)
	copy(e.S_free_blocks_count[:], fcontBlocks)
	copy(e.S_free_inodes_count[:], fcontInodos)
	copy(e.S_mtime[:], mtime)
	copy(e.S_mnt_count[:], mnt_cont)
	copy(e.S_magic[:], magic)
	copy(e.S_inode_size[:], inodeTam)
	copy(e.S_block_size[:], blockTam)
	copy(e.S_first_ino[:], primIno)
	copy(e.S_first_blo[:], primBlock)
	copy(e.S_bm_inode_start[:], bm_inodest)
	copy(e.S_bm_block_start[:], bm_blockst)
	copy(e.S_inode_start[:], inodest)
	copy(e.S_block_start[:], blockst)
	return e
}

type TablaInodos struct{
	I_uid [10000]byte
	I_gid [10000]byte
	I_size [10000]byte
	I_atime [10000]byte
	I_ctime [10000]byte
	I_mtime [10000]byte
	I_block [15]int
	I_type [10000]byte
	I_perm [10000]byte
}

func NewTablaInodos(uid, gid, size, atime, ctime, mtime, typ, perm string) TablaInodos{
	e := TablaInodos{}
	copy(e.I_uid[:],uid) 
	copy(e.I_gid[:],gid)
	copy(e.I_size[:],size) 
	copy(e.I_atime[:],atime) 
	copy(e.I_ctime[:],ctime) 
	copy(e.I_mtime[:],mtime) 
	for i := 0; i < 15; i++{
		e.I_block[i] = -1
	}
	copy(e.I_type[:],typ) 
	copy(e.I_perm[:],perm)
	return e 
}


type BlockDirectory struct{
	B_content [4]Content
}


type Content struct{
	B_name [100]byte
	B_inodo [100]byte
}

func NewContent(nom, ino string)Content{
	e := Content{}
	copy(e.B_name[:], nom)
	copy(e.B_inodo[:], ino)
	return e
}

type BlockFile struct{
	B_content [64]byte
}

func NewBlockFile(cont string)BlockFile{
	e := BlockFile{}
	copy(e.B_content[:], cont)
	return e
}

type ParticionesMontadas struct{
	Name string
	Id string
	Inicio int
}

type DiscosMontados struct{
	Path string
	Letra string
	Particiones_montadas [60]ParticionesMontadas
}

type Grupo struct{
	Id string
	Name string
}

type Usuario struct{
	Id string
	Tipo string
	Name string
	Pwd string
	Group string
}

type LoginSesion struct{
	User string
	Path string
	NomPart string
	Id string
	Id_usuario string
	Id_grupo string 
}
