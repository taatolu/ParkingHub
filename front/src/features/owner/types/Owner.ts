export type Owner = {
    id: number;
    first_name: string;
    middle_name: string;
    last_name: string;
    license_expiration: string;
};

// 新規オーナー作成用の型（idを除く）
export type CreateOwnerDTO = Omit<Owner, 'id'>;