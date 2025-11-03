import Reacr, { Component } from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { OwnerDashboard } from './OwnerDashboard';
import { useOwnersList } from '../features/owner/hooks/useOwnersList';
import { useOwnerDetails } from '../features/owner/hooks/useOwnersDetail';
import { Owner } from '../features/owner/types/Owner';

// モックの設定
jest.mock('../features/owner/hooks/useOwnersList');
jest.mock('../features/owner/hooks/useOwnersDetail');

//コンポーネントのモック
/// OwnerListコンポーネントのモック
jest.mock('../features/owner/components/OwnerList', () => (
    { OwnerList: ({ owners, onDetail, onDelete }: {owners:Owner[]; onDetail: (id:number) => void; onDelete: (id:number)=> void})  => (
        <div data-testid="owner-list-mock">
            {owners.map(owner => (
                <div key={owner.id}>
                    <div>{owner.first_name}</div>
                    <div>
                        <button onClick={() => onDetail(owner.id)}>詳細</button>
                        <button onClick={() => onDelete(owner.id)}>削除</button>
                    </div>
                </div>
            ))}
        </div>)}));

/// EditOwnerModalコンポーネントのモック
jest.mock('../features/owner/components/EditOwnerModal', () => (
    { EditOwnerModal: ({ isOpen, onClose}: { isOpen: boolean; onClose: () => void}) => (
        isOpen ? <div data-testid="edit-owner-modal-mock">
            <button onClick={onClose}>閉じる</button>
        </div> : null
    )}
));

// テストデータの定義
const mockOwners: Owner[] = [
    {id: 1, first_name: 'Taro', middle_name: '', last_name: 'Yamada', license_expiration: '2023-12-31'},
    {id: 2, first_name: 'Hanako', middle_name: '', last_name: 'Suzuki', license_expiration: '2022-06-30'},
];

// テストケース
describe('OwnerDashboard', () => {
    beforeEach(() => {
        // useOwnersListのモックの初期化
        (useOwnersList as jest.Mock).mockReturnValue({
            owners: mockOwners,
            loading: false,
            error: null,
            fetchOwners: jest.fn(),
            showExpiredOnly: false,
            toggleExpiredFilter: jest.fn(),
        });

        // useOwnerDetailsのモックの初期化
        (useOwnerDetails as jest.Mock).mockReturnValue({
            selectedOwner: null,
            loading: false,
            error: null,
            fetchOwnerDetails: jest.fn(),
            createOwnerDetails: jest.fn(),
            updateOwnerDetails: jest.fn(),
        });
    });

    test('オーナー一覧が正しく表示されること', () => {
        render(<OwnerDashboard />);     // コンポーネントのレンダリング
        expect(screen.getByText('オーナー一覧')).toBeInTheDocument();   // タイトルの表示があるか確認
        expect(screen.getByText('Taro')).toBeInTheDocument();   // 初期化で渡したオーナーデータが表示されているか確認
        expect(screen.getByText('Hanako')).toBeInTheDocument();  // 初期化で渡したオーナーデータが表示されているか確認
    });

    test('詳細ボタンをクリックするとEditOwnerModalが表示されること', async () => {
        const { fetchOwnerDetails } = useOwnerDetails();    // フックからfetchOwnerDetails関数を取得
        render(<OwnerDashboard />); // Dashboardコンポーネントのレンダリング
        
        // 詳細ボタンをクリックしてモーダルを開く(0番目のオーナー)
        fireEvent.click(screen.getAllByText('詳細')[0]);

        // awaitで非同期処理を待ち、モーダルが表示されることを確認
        await waitFor(() => {
            expect(fetchOwnerDetails).toHaveBeenCalledWith(1); // fetchOwnerDetailsが正しいIDで呼ばれたか確認
            expect(screen.getByTestId('edit-owner-modal-mock')).toBeInTheDocument();    // モーダルが表示されていることを確認
        });
    });

    test('モーダルの閉じるボタンをクリックするとEditOwnerModalが閉じること', async () => {
        render(<OwnerDashboard />);
        
        // 詳細ボタンをクリックしてモーダルを開く
        fireEvent.click(screen.getAllByText('詳細')[0]);
        
        await waitFor(() => {
            expect(screen.getByTestId('edit-owner-modal-mock')).toBeInTheDocument();
        });

        // 閉じるボタンをクリックしてモーダルを閉じる
        fireEvent.click(screen.getByText('閉じる'));
        
        await waitFor(() => {
            expect(screen.queryByTestId('edit-owner-modal-mock')).not.toBeInTheDocument();
        });
    });
});